<script>
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import Chart from "chart.js/auto";
    import Navbar from "../../../components/Navbar.svelte";

    let electionID;
    let voteData = [];
    let chartInstances = {};
    let candidatesMap = {};
    let electionInfo = null;
    let showResults = false;

    let selectedPost = "";
    let selectedParty = "";

    let posts = new Set();
    let parties = new Set();

    $: electionID = $page.params.id;

    $: filteredCandidates = Object.values(candidatesMap).filter((c) => {
        return (
            (selectedPost ? c.candidate_post === selectedPost : true) &&
            (selectedParty ? c.candidate_party === selectedParty : true)
        );
    });

    async function fetchElectionInfo() {
        const token = localStorage.getItem("user_token");
        const res = await fetch(
            `http://localhost:3000/election/${electionID}`,
            {
                headers: { Authorization: `Bearer ${token}` },
            },
        );
        if (!res.ok) throw new Error("Could not fetch election details");
        electionInfo = await res.json();

        const now = new Date();
        const endDate = new Date(electionInfo.end_date);
        showResults = now >= endDate;
    }

    async function fetchCandidates() {
        const token = localStorage.getItem("user_token");
        if (!token) return;

        try {
            const res = await fetch(
                `http://localhost:3000/candidate?limit=10&offset=0`,
                {
                    headers: { Authorization: `Bearer ${token}` },
                },
            );
            if (res.ok) {
                const candidates = await res.json();
                candidates.forEach((c) => {
                    candidatesMap[c.candidate_id] = c;
                    posts.add(c.candidate_post);
                    parties.add(c.candidate_party);
                });
            } else {
                console.warn("Failed to fetch candidates.");
            }
        } catch (error) {
            console.error("Error fetching candidates:", error);
        }
    }

    async function fetchData() {
        const token = localStorage.getItem("user_token");
        if (!token) return goto("/");
        if (!electionID) return console.error("Election ID is missing!");

        try {
            const res = await fetch(
                `http://localhost:3000/vote/history/${electionID}?limit=1000&offset=0`,
                { headers: { Authorization: `Bearer ${token}` } },
            );
            if (!res.ok) throw new Error("Failed to fetch votes");

            voteData = await res.json();
            await fetchCandidates();
            createGraphs();
        } catch (error) {
            console.error(error);
        }
    }

    onMount(async () => {
        try {
            await fetchElectionInfo();
            if (showResults) await fetchData();
        } catch (err) {
            console.error("Error:", err);
        }
    });

    function destroyCharts() {
        for (const chartId in chartInstances) {
            chartInstances[chartId]?.destroy();
        }
        chartInstances = {};
    }

    function createGraphs() {
        destroyCharts();
        const voteCountsByPost = {};

        voteData.forEach((vote) => {
            const candidate = candidatesMap[vote.candidate_id];
            if (!candidate) return;
            const post = candidate.candidate_post || "Unknown";

            if (!voteCountsByPost[post]) voteCountsByPost[post] = {};
            voteCountsByPost[post][candidate.candidate_id] =
                (voteCountsByPost[post][candidate.candidate_id] || 0) + 1;
        });

        Object.entries(voteCountsByPost).forEach(([post, counts]) => {
            const canvasId = `chart-${post.replace(/\s+/g, "-").toLowerCase()}`;
            const labels = Object.keys(counts).map(
                (id) => candidatesMap[id].candidate_name,
            );
            const data = Object.values(counts);

            createBarChart(
                canvasId,
                `${post} Vote Distribution`,
                labels,
                data,
                "Votes",
            );
        });

        const totalVotes = voteData.length;
        const votesPerCandidate = {};

        voteData.forEach((vote) => {
            votesPerCandidate[vote.candidate_id] =
                (votesPerCandidate[vote.candidate_id] || 0) + 1;
        });

        Object.entries(votesPerCandidate).forEach(([id, count]) => {
            const candidate = candidatesMap[id];
            const canvasId = `pie-${id}`;
            const name = candidate?.candidate_name || `Candidate ${id}`;
            createPieChart(
                canvasId,
                `${name} Vote Share`,
                [name, "Others"],
                [count, totalVotes - count],
            );
        });
    }

    function createBarChart(canvasId, title, labels, data, yLabel) {
        const ctx = document.getElementById(canvasId);
        if (ctx) {
            chartInstances[canvasId] = new Chart(ctx, {
                type: "bar",
                data: {
                    labels,
                    datasets: [
                        {
                            label: title,
                            data,
                            backgroundColor: "rgba(54, 162, 235, 0.8)",
                            borderColor: "rgba(54, 162, 235, 1)",
                            borderWidth: 1,
                        },
                    ],
                },
                options: {
                    responsive: true,
                    scales: {
                        y: {
                            beginAtZero: true,
                            title: { display: true, text: yLabel },
                        },
                    },
                    plugins: {
                        title: { display: true, text: title },
                        legend: { display: false },
                    },
                },
            });
        }
    }

    function createPieChart(canvasId, title, labels, data) {
        const ctx = document.getElementById(canvasId);
        if (ctx) {
            chartInstances[canvasId] = new Chart(ctx, {
                type: "pie",
                data: {
                    labels,
                    datasets: [
                        {
                            data,
                            backgroundColor: [
                                "rgba(255, 99, 132, 0.8)",
                                "rgba(54, 162, 235, 0.8)",
                            ],
                            borderColor: "rgba(255, 255, 255, 1)",
                            borderWidth: 1,
                        },
                    ],
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: { display: true, text: title },
                        legend: { position: "bottom" },
                    },
                },
            });
        }
    }
</script>

<svelte:head>
    <title>Election Results</title>
</svelte:head>

<Navbar />

<div class="max-w-5xl mx-auto p-4">
    <h1 class="text-3xl font-bold mb-4">Election Results (ID: {electionID})</h1>

    {#if !showResults}
        <div class="alert alert-info shadow-lg mt-10">
            <div>
                <span>
                    This election hasn't ended yet. Results are still being
                    processed.
                </span>
            </div>
        </div>
    {:else}
        <div class="flex flex-col gap-4 md:flex-row mb-6">
            <select
                class="select select-bordered w-full"
                bind:value={selectedPost}
            >
                <option value="">All Posts</option>
                {#each Array.from(posts) as post}
                    <option value={post}>{post}</option>
                {/each}
            </select>

            <select
                class="select select-bordered w-full"
                bind:value={selectedParty}
            >
                <option value="">All Parties</option>
                {#each Array.from(parties) as party}
                    <option value={party}>{party}</option>
                {/each}
            </select>
        </div>

        {#each filteredCandidates as candidate (candidate.candidate_id)}
            <div class="bg-white shadow-md rounded-xl p-4 mb-8 border">
                <h3 class="text-lg font-semibold mb-2">
                    {candidate.candidate_post} - {candidate.candidate_name}
                </h3>
                <p class="text-sm mb-2 text-gray-600">
                    Party: {candidate.candidate_party}
                </p>
                <canvas
                    id={"chart-" +
                        candidate.candidate_post
                            .replace(/\s+/g, "-")
                            .toLowerCase()}
                ></canvas>
                <div class="mt-6">
                    <canvas id={"pie-" + candidate.candidate_id}></canvas>
                </div>
            </div>
        {/each}
    {/if}
</div>
