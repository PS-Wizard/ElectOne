<script>
    import { onMount, onDestroy } from "svelte";
    import { fade } from "svelte/transition";
    import { page } from "$app/stores";
    import {
        Chart,
        BarController,
        CategoryScale,
        LinearScale,
        BarElement,
        PieController,
        ArcElement,
        Tooltip,
        Legend,
    } from "chart.js";
    import Navbar from "../../../components/Navbar.svelte";
    import CandidatePieChartCard from "../../../components/CandidatePieChartCard.svelte";

    Chart.register(
        BarController,
        CategoryScale,
        LinearScale,
        BarElement,
        PieController,
        ArcElement,
        Tooltip,
        Legend,
    );

    const API_BASE_URL = "http://localhost:3000";

    $: electionId = $page.params.id;
    let election = {
        election_id: electionId,
        name: "Loading...",
        description: "",
        start_date: "",
        end_date: "",
        location: "",
        type: "",
    };

    let candidates = [];
    let votes = [];
    let loading = true;
    let error = null;

    let postsWithCandidatesAndVotes = [];
    let barChartInstances = {};
    let ws;

    onMount(async () => {
        try {
            const electionRes = await fetch(
                `${API_BASE_URL}/election/${electionId}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("user_token")}`,
                    },
                },
            );

            if (!electionRes.ok)
                throw new Error(
                    `Failed to fetch election: ${electionRes.statusText}`,
                );
            election = await electionRes.json();

            const candidatesRes = await fetch(
                `${API_BASE_URL}/election/candidates/${electionId}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("user_token")}`,
                    },
                },
            );
            if (!candidatesRes.ok)
                throw new Error(
                    `Failed to fetch candidates: ${candidatesRes.statusText}`,
                );
            candidates = await candidatesRes.json();

            const now = new Date();
            const electionEnded = election.end_date
                ? new Date(election.end_date) < now
                : false;

            if (electionEnded) {
                await fetchVotesAndRender();
            } else {
                await fetchVotesAndRender();
                setupWebSocket();
            }
        } catch (err) {
            error = err.message;
            console.error("Error fetching data:", err);
        } finally {
            loading = false;
        }
    });

    onDestroy(() => {
        Object.values(barChartInstances).forEach((chart) => {
            if (chart) chart.destroy();
        });
        barChartInstances = {};

        if (ws) {
            ws.close();
        }
    });

    async function fetchVotesAndRender() {
        const electionEnded = election.end_date
            ? new Date(election.end_date) < new Date()
            : false;
        const endpoint = electionEnded
            ? `${API_BASE_URL}/vote/turso/history/${electionId}`
            : `${API_BASE_URL}/vote/redis/history/${electionId}`;

        const votesRes = await fetch(`${endpoint}?limit=100&offset=0`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem("user_token")}`,
            },
        });

        if (!votesRes.ok)
            throw new Error(`Failed to fetch votes: ${votesRes.statusText}`);
        votes = await votesRes.json();

        processDataAndRenderBarCharts();
    }

    function setupWebSocket() {
        ws = new WebSocket("ws://localhost:3000/live/votes");

        ws.onmessage = (event) => {
            try {
                const newVote = JSON.parse(event.data);
                console.log(newVote);

                votes.push(newVote);
                processDataAndRenderBarCharts();
            } catch (err) {
                console.error("Invalid WS message", err);
            }
        };

        ws.onerror = (err) => {
            console.error("WebSocket error", err);
        };

        ws.onclose = () => {
            console.log("WebSocket closed");
        };
    }

    function processDataAndRenderBarCharts() {
        Object.values(barChartInstances).forEach((chart) => {
            if (chart) chart.destroy();
        });
        barChartInstances = {};
        const voteCounts = {};
        votes.forEach((vote) => {
            voteCounts[vote.candidate_id] =
                (voteCounts[vote.candidate_id] || 0) + 1;
        });

        const candidatesWithVotes = candidates.map((candidate) => ({
            ...candidate,
            votes: voteCounts[candidate.candidate_id] || 0,
        }));

        const uniquePostNames = [
            ...new Set(candidatesWithVotes.map((c) => c.candidate_post)),
        ];

        postsWithCandidatesAndVotes = uniquePostNames.map((postName) => {
            const postCandidates = candidatesWithVotes.filter(
                (c) => c.candidate_post === postName,
            );
            const totalVotesInPost = postCandidates.reduce(
                (sum, c) => sum + c.votes,
                0,
            );

            postCandidates.sort((a, b) => b.votes - a.votes);

            return {
                name: postName,
                candidates: postCandidates,
                totalVotesInPost,
            };
        });

        setTimeout(() => {
            postsWithCandidatesAndVotes.forEach((postData) => {
                renderBarChart(postData.name, postData.candidates);
            });
        }, 100);
    }

    function renderBarChart(postName, postCandidates) {
        const winner = postCandidates.length > 0 ? postCandidates[0] : null;
        const canvasId = `bar-chart-${postName.replace(/\s+/g, "-")}`;
        const ctx = document.getElementById(canvasId);
        if (!ctx) {
            console.warn(`Canvas element not found for bar chart: ${canvasId}`);
            return;
        }

        if (barChartInstances[postName]) {
            barChartInstances[postName].destroy();
        }

        const chart = new Chart(ctx, {
            type: "bar",
            data: {
                labels: postCandidates.map((c) => c.candidate_name),
                datasets: [
                    {
                        label: "Votes",
                        data: postCandidates.map((c) => c.votes),
                        backgroundColor: postCandidates.map((c) =>
                            winner && c.candidate_id === winner.candidate_id
                                ? "rgba(255, 215, 0, 0.7)"
                                : "rgba(75, 192, 192, 0.7)",
                        ),
                        borderColor: postCandidates.map((c) =>
                            winner && c.candidate_id === winner.candidate_id
                                ? "rgb(255, 215, 0)"
                                : "rgb(75, 192, 192)",
                        ),
                        borderWidth: 1,
                    },
                ],
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    title: {
                        display: true,
                        text: `Vote Results for ${postName}`,
                        font: { size: 16 },
                    },
                    tooltip: {
                        callbacks: {
                            label: (context) => `Votes: ${context.raw}`,
                        },
                    },
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        title: { display: true, text: "Number of Votes" },
                    },
                    x: {
                        title: { display: true, text: "Candidates" },
                    },
                },
            },
        });

        barChartInstances[postName] = chart;
    }
</script>

<svelte:head>
    <title
        >Election Results - {election.name === "Loading..."
            ? "Loading"
            : election.name}</title
    >
</svelte:head>

<div
    class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
    aria-hidden="true"
>
    <div
        class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
        style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
    ></div>
</div>

<Navbar />

<div class="container mx-auto px-4 py-8">
    {#if loading}
        <div class="text-center py-8">
            <div
                class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500 mx-auto"
            ></div>
            <p class="mt-4 text-gray-600">Loading election data...</p>
        </div>
    {:else if error}
        <div
            class="bg-red-100 border-l-4 border-red-500 text-red-700 p-4 mb-4"
            role="alert"
        >
            <p class="font-bold">Error</p>
            <p>{error}</p>
        </div>
    {:else}
        <h1 class="text-3xl font-bold uppercase text-center mb-6">
            {election.name} Results
        </h1>

        <div class="bg-white rounded-lg p-6 mb-8 border">
            <h2 class="text-xl font-semibold mb-4">Election Details</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                    <p>
                        <span class="font-medium">Location:</span>
                        {election.location}
                    </p>
                    <p>
                        <span class="font-medium">Type:</span>
                        {election.type}
                    </p>
                </div>
                <div>
                    <p>
                        <span class="font-medium">Dates:</span>
                        {election.start_date
                            ? new Date(election.start_date).toLocaleDateString()
                            : "N/A"} to
                        {election.end_date
                            ? new Date(election.end_date).toLocaleDateString()
                            : "N/A"}
                    </p>
                    <p>
                        <span class="font-medium">Total Votes Cast:</span>
                        {votes.length}
                    </p>
                </div>
            </div>
            {#if election.description}
                <p class="mt-4">{election.description}</p>
            {/if}
        </div>

        {#each postsWithCandidatesAndVotes as postData (postData.name)}
            <div class="bg-white rounded-lg p-6 mb-8 border" in:fade out:fade>
                <h2 class="text-xl font-semibold mb-4 uppercase tracking-wide">
                    {postData.name}
                </h2>
                <div class="chart-container relative h-64">
                    <canvas
                        id={"bar-chart-" + postData.name.replace(/\s+/g, "-")}
                    ></canvas>
                </div>

                <div class="mt-6 grid grid-cols-1 md:grid-cols-3 gap-6">
                    {#each postData.candidates as candidate (candidate.candidate_id)}
                        <CandidatePieChartCard
                            {candidate}
                            election_id={election.election_id}
                            totalVotesInPost={postData.totalVotesInPost}
                        />
                    {/each}
                </div>
            </div>
        {/each}
    {/if}
</div>

<style>
    .chart-container canvas {
        max-width: 100% !important;
        max-height: 400px !important;
    }
</style>
