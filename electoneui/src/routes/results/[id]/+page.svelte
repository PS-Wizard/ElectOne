<script>
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import Chart from "chart.js/auto";
    import Navbar from "../../../components/Navbar.svelte";

    let electionID;
    let voteData = [];
    let chartInstances = {};
    let candidatesMap = {}; // candidate_id -> full candidate obj

    $: electionID = $page.params.id;

    async function fetchCandidates() {
        const token = localStorage.getItem("user_token");
        if (!token) {
            error = "You're not logged in.";
            loading = false;
            return;
        }
        try {
            const res = await fetch(
                `http://localhost:3000/election/candidates/${electionID}`,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                },
            );
            if (res.ok) {
                const candidates = await res.json();
                candidates.forEach((c) => {
                    candidatesMap[c.candidate_id] = c;
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
        if (!token) {
            goto("/");
            return;
        }
        if (!electionID) {
            console.error("Election ID is missing!");
            return;
        }

        console.log("Attempting to fetch data for election:", electionID);

        try {
            const res = await fetch(
                `http://localhost:3000/vote/history/${electionID}?limit=1000&offset=0`,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                },
            );
            if (!res.ok) {
                throw new Error("Failed to fetch votes");
            }
            voteData = await res.json();
            console.log("Vote Data:", voteData);
            await fetchCandidates();
            createGraphs();
        } catch (error) {
            console.error(error);
        }
    }

    onMount(() => {
        fetchData();
        //const socket = new WebSocket("ws://localhost:3000/live/ws/votes");

        //socket.onmessage = async (event) => {
        //    const vote = JSON.parse(event.data);
        //    if (vote.election_id == electionID) {
        //        voteData = [...voteData, vote];
        //        createGraphs(); // refresh charts
        //    }
        //};

        //socket.onerror = (e) => console.error("WebSocket error:", e);

        //return () => {
        //    socket.close();
        //};
    });

    function destroyCharts() {
        for (const chartId in chartInstances) {
            if (chartInstances[chartId]) {
                chartInstances[chartId].destroy();
            }
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

        // 🔥 One bar chart per post
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

        // 🔥 Pie chart per candidate (their vote %)
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
                    labels: labels,
                    datasets: [
                        {
                            label: title,
                            data: data,
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
                            title: {
                                display: true,
                                text: yLabel,
                            },
                        },
                    },
                    plugins: {
                        title: {
                            display: true,
                            text: title,
                        },
                        legend: {
                            display: false,
                        },
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
                    labels: labels,
                    datasets: [
                        {
                            data: data,
                            backgroundColor: [
                                "rgba(255, 99, 132, 0.8)",
                                "rgba(54, 162, 235, 0.8)",
                                "rgba(255, 206, 86, 0.8)",
                                "rgba(75, 192, 192, 0.8)",
                                "rgba(153, 102, 255, 0.8)",
                                "rgba(255, 159, 64, 0.8)",
                            ],
                            borderColor: "rgba(255, 255, 255, 1)",
                            borderWidth: 1,
                        },
                    ],
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: title,
                        },
                        legend: {
                            position: "bottom",
                        },
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
<h1>Election Results for ID: {electionID}</h1>

{#each Object.values(candidatesMap) as candidate}
    <div class="chart-container">
        <h3>{candidate.candidate_post} - {candidate.candidate_name}</h3>
        <canvas
            id={"chart-" +
                candidate.candidate_post.replace(/\s+/g, "-").toLowerCase()}
        ></canvas>
    </div>
{/each}

{#each Object.keys(candidatesMap) as id}
    <div class="chart-container">
        <h3>Vote Share for {candidatesMap[id].candidate_name}</h3>
        <canvas id={"pie-" + id}></canvas>
    </div>
{/each}

<style>
    .chart-container {
        width: 80%;
        margin: 20px auto;
        border: 1px solid #ccc;
        padding: 10px;
        border-radius: 8px;
    }

    canvas {
        max-width: 100%;
        height: auto;
    }
</style>
