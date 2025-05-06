<script>
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import Chart from "chart.js/auto";

    let electionID;
    let voteData = [];
    let candidateNames = {};
    let chartInstances = {};

    $: electionID = $page.params.id;

    async function fetchCandidateNames() {
        try {
            const res = await fetch(
                `http://localhost:3000/elections/${electionID}/candidates`,
            );
            if (res.ok) {
                const candidates = await res.json();
                candidates.forEach((candidate) => {
                    candidateNames[candidate.candidate_id] =
                        candidate.name || `Candidate ${candidate.candidate_id}`;
                });
            } else {
                console.warn("Failed to fetch candidate names.");
                const uniqueCandidateIds = [
                    ...new Set(voteData.map((vote) => vote.candidate_id)),
                ];
                uniqueCandidateIds.forEach(
                    (id) => (candidateNames[id] = `Candidate ${id}`),
                );
            }
        } catch (error) {
            console.error("Error fetching candidate names:", error);
            const uniqueCandidateIds = [
                ...new Set(voteData.map((vote) => vote.candidate_id)),
            ];
            uniqueCandidateIds.forEach(
                (id) => (candidateNames[id] = `Candidate ${id}`),
            );
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
            await fetchCandidateNames();
            createGraphs();
        } catch (error) {
            console.error(error);
        }
    }

    onMount(() => {
        fetchData();
        const socket = new WebSocket("ws://localhost:3000/live/ws/votes");

        socket.onmessage = async (event) => {
            const vote = JSON.parse(event.data);
            if (vote.election_id == electionID) {
                voteData = [...voteData, vote];
                createGraphs(); // refresh charts
            }
        };

        socket.onerror = (e) => console.error("WebSocket error:", e);

        return () => {
            socket.close();
        };
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

        // 1. Candidate Vote Distribution
        const candidateVoteCounts = {};
        voteData.forEach((vote) => {
            candidateVoteCounts[vote.candidate_id] =
                (candidateVoteCounts[vote.candidate_id] || 0) + 1;
        });

        const candidateLabels = Object.keys(candidateVoteCounts).map(
            (id) => candidateNames[id],
        );
        const voteCounts = Object.values(candidateVoteCounts);
        createBarChart(
            "candidateVoteChart",
            "Candidate Vote Distribution",
            candidateLabels,
            voteCounts,
            "Number of Votes",
        );

        // 2. Candidate Vote Percentage
        const totalVotes = voteData.length;
        const candidatePercentages = {};
        for (const candidateId in candidateVoteCounts) {
            candidatePercentages[candidateId] =
                (candidateVoteCounts[candidateId] / totalVotes) * 100;
        }

        const percentageLabels = Object.keys(candidatePercentages).map(
            (id) => candidateNames[id],
        );
        const percentageValues = Object.values(candidatePercentages);
        createPieChart(
            "candidatePercentageChart",
            "Candidate Vote Percentage",
            percentageLabels,
            percentageValues,
        );

        // 3. vote count by candidate id
        const candidateIdCounts = {};
        voteData.forEach((vote) => {
            candidateIdCounts[vote.candidate_id] =
                (candidateIdCounts[vote.candidate_id] || 0) + 1;
        });
        const candidateIdLabels = Object.keys(candidateIdCounts);
        const candidateIdValues = Object.values(candidateIdCounts);
        createBarChart(
            "candidateIdChart",
            "Vote Count by Candidate ID",
            candidateIdLabels,
            candidateIdValues,
            "Vote Count",
        );
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

<h1>Election Results for ID: {electionID}</h1>

<div class="chart-container">
    <canvas id="candidateVoteChart"></canvas>
</div>

<div class="chart-container">
    <canvas id="candidatePercentageChart"></canvas>
</div>
<div class="chart-container">
    <canvas id="candidateIdChart"></canvas>
</div>

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
