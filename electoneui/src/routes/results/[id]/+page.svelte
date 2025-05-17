<script>
    import { onMount } from "svelte";
    import { fade } from "svelte/transition";
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

    // Register Chart.js components
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

    // Sample data structure
    let election = {
        election_id: 1,
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
    let posts = [];
    let charts = []; // To store chart instances for cleanup

    // Fetch data on mount
    onMount(async () => {
        try {
            // Fetch election details
            const electionRes = await fetch(`${API_BASE_URL}/election/1`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("user_token")}`,
                },
            });
            election = await electionRes.json();

            // Fetch candidates
            const candidatesRes = await fetch(
                `${API_BASE_URL}/election/candidates/1`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("user_token")}`,
                    },
                },
            );
            candidates = await candidatesRes.json();

            // Fetch votes
            const votesRes = await fetch(
                `${API_BASE_URL}/vote/history/1?limit=100&offset=0`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("user_token")}`,
                    },
                },
            );
            votes = await votesRes.json();

            // Process data
            processData();
        } catch (err) {
            error = err.message;
            console.error("Error fetching data:", err);
        } finally {
            loading = false;
        }

        // Cleanup charts when component is destroyed
        return () => {
            charts.forEach((chart) => chart.destroy());
        };
    });

    // Process data
    function processData() {
        // Group votes by candidate_id
        const voteCounts = {};
        votes.forEach((vote) => {
            voteCounts[vote.candidate_id] =
                (voteCounts[vote.candidate_id] || 0) + 1;
        });

        // Add vote counts to candidates
        candidates.forEach((candidate) => {
            candidate.votes = voteCounts[candidate.candidate_id] || 0;
        });

        // Find unique posts
        posts = [...new Set(candidates.map((c) => c.candidate_post))];

        // Render charts after a small delay to ensure DOM is ready
        setTimeout(renderCharts, 100);
    }

    // Render all charts
    function renderCharts() {
        posts.forEach((post) => {
            const postCandidates = candidates.filter(
                (c) => c.candidate_post === post,
            );
            renderBarChart(post, postCandidates);
            renderPieCharts(post, postCandidates);
        });
    }

    // Render bar chart for a specific post
    function renderBarChart(post, postCandidates) {
        // Sort by votes descending
        postCandidates.sort((a, b) => b.votes - a.votes);

        // Find the winner
        const winner = postCandidates[0];

        const ctx = document.getElementById(
            `bar-chart-${post.replace(/\s+/g, "-")}`,
        );
        if (!ctx) return;

        // Destroy previous chart if it exists
        const existingChart = Chart.getChart(ctx);
        if (existingChart) existingChart.destroy();

        const chart = new Chart(ctx, {
            type: "bar",
            data: {
                labels: postCandidates.map((c) => c.candidate_name),
                datasets: [
                    {
                        label: "Votes",
                        data: postCandidates.map((c) => c.votes),
                        backgroundColor: postCandidates.map((c) =>
                            c.candidate_id === winner.candidate_id
                                ? "rgba(255, 215, 0, 0.7)"
                                : "rgba(75, 192, 192, 0.7)",
                        ),
                        borderColor: postCandidates.map((c) =>
                            c.candidate_id === winner.candidate_id
                                ? "rgb(255, 215, 0)"
                                : "rgb(75, 192, 192)",
                        ),
                        borderWidth: 1,
                    },
                ],
            },
            options: {
                responsive: true,
                plugins: {
                    title: {
                        display: true,
                        text: `Vote Results for ${post}`,
                        font: {
                            size: 16,
                        },
                    },
                    tooltip: {
                        callbacks: {
                            label: function (context) {
                                return `Votes: ${context.raw}`;
                            },
                        },
                    },
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: "Number of Votes",
                        },
                    },
                    x: {
                        title: {
                            display: true,
                            text: "Candidates",
                        },
                    },
                },
            },
        });

        charts.push(chart);
    }

    // Render pie charts for each candidate in a post
    function renderPieCharts(post, postCandidates) {
        const container = document.getElementById(
            `pie-charts-${post.replace(/\s+/g, "-")}`,
        );
        if (!container) return;

        // Clear previous content
        container.innerHTML = "";

        // Calculate total votes for this post
        const totalVotes = postCandidates.reduce((sum, c) => sum + c.votes, 0);

        // Create a chart for each candidate
        postCandidates.forEach((candidate) => {
            const percentage =
                totalVotes > 0
                    ? ((candidate.votes / totalVotes) * 100).toFixed(1)
                    : 0;

            const card = document.createElement("div");
            card.className = "candidate-card";
            card.style.width = "250px";
            card.style.margin = "10px";
            card.style.padding = "15px";
            card.style.borderRadius = "8px";
            card.style.boxShadow = "0 2px 8px rgba(0,0,0,0.1)";
            card.style.textAlign = "center";

            // Candidate image
            const img = document.createElement("img");
            img.src = `${candidate.candidate_photo}`;
            img.alt = candidate.candidate_name;
            img.style.width = "80px";
            img.style.height = "80px";
            img.style.borderRadius = "50%";
            img.style.objectFit = "cover";
            img.style.marginBottom = "10px";
            card.appendChild(img);

            // Candidate name
            const name = document.createElement("h3");
            name.textContent = candidate.candidate_name;
            name.style.margin = "5px 0";
            name.style.fontSize = "16px";
            card.appendChild(name);

            // Candidate party
            const party = document.createElement("p");
            party.textContent = candidate.candidate_party;
            party.style.margin = "5px 0";
            party.style.fontSize = "14px";
            party.style.color = "#666";
            card.appendChild(party);

            // Vote percentage
            const percent = document.createElement("div");
            percent.textContent = `${percentage}% of votes`;
            percent.style.margin = "10px 0";
            percent.style.fontSize = "18px";
            percent.style.fontWeight = "bold";
            card.appendChild(percent);

            // Create canvas for pie chart
            const canvas = document.createElement("canvas");
            canvas.width = 200;
            canvas.height = 200;
            card.appendChild(canvas);

            container.appendChild(card);

            // Create pie chart
            const ctx = canvas.getContext("2d");
            const chart = new Chart(ctx, {
                type: "pie",
                data: {
                    labels: ["This Candidate", "Others"],
                    datasets: [
                        {
                            data: [
                                candidate.votes,
                                totalVotes - candidate.votes,
                            ],
                            backgroundColor: [
                                "rgba(75, 192, 192, 0.7)",
                                "rgba(201, 203, 207, 0.7)",
                            ],
                            borderColor: [
                                "rgba(75, 192, 192, 1)",
                                "rgba(201, 203, 207, 1)",
                            ],
                            borderWidth: 1,
                        },
                    ],
                },
                options: {
                    responsive: false,
                    plugins: {
                        legend: {
                            position: "bottom",
                        },
                        tooltip: {
                            callbacks: {
                                label: function (context) {
                                    const label = context.label || "";
                                    const value = context.raw || 0;
                                    const total = context.dataset.data.reduce(
                                        (a, b) => a + b,
                                        0,
                                    );
                                    const percentage = Math.round(
                                        (value / total) * 100,
                                    );
                                    return `${label}: ${value} (${percentage}%)`;
                                },
                            },
                        },
                    },
                },
            });

            charts.push(chart);
        });
    }
</script>

<svelte:head>
    <title>Election Results - {election.name}</title>
</svelte:head>

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
        <h1 class="text-3xl font-bold text-center mb-6">
            {election.name} Results
        </h1>

        <div class="bg-white rounded-lg shadow-md p-6 mb-8">
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
                        {new Date(election.start_date).toLocaleDateString()} to {new Date(
                            election.end_date,
                        ).toLocaleDateString()}
                    </p>
                    <p>
                        <span class="font-medium">Total Votes:</span>
                        {votes.length}
                    </p>
                </div>
            </div>
            <p class="mt-4">
                <span class="font-medium">Description:</span>
                {election.description}
            </p>
        </div>

        {#each posts as post}
            <div class="bg-white rounded-lg shadow-md p-6 mb-8" transition:fade>
                <h2 class="text-2xl font-bold mb-6 text-center">
                    {post} Results
                </h2>

                <!-- Bar Chart -->
                <div class="mb-8">
                    <h3 class="text-lg font-semibold mb-4">
                        Vote Distribution
                    </h3>
                    <div class="w-full h-96">
                        <canvas id="bar-chart-{post.replace(/\s+/g, '-')}"
                        ></canvas>
                    </div>
                </div>

                <!-- Pie Charts -->
                <div class="mb-8">
                    <h3 class="text-lg font-semibold mb-4">
                        Candidate Performance
                    </h3>
                    <div
                        id="pie-charts-{post.replace(/\s+/g, '-')}"
                        class="flex flex-wrap justify-center gap-6"
                    ></div>
                </div>
            </div>
        {/each}
    {/if}
</div>

<style>
    .candidate-card {
        transition:
            transform 0.2s ease,
            box-shadow 0.2s ease;
    }

    .candidate-card:hover {
        transform: translateY(-5px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    }
</style>
