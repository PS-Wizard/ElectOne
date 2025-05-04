<script>
    import { onMount } from "svelte";
    import Chart from "chart.js/auto";
    import Navbar from "../../../../../components/Navbar.svelte";

    let electionId = 5;
    let historyData = {};
    let chartInstances = {};
    let socket;
    let selectedPost;
    let canvasElements = {}; // Store canvas elements by ID
    let initialRender = true;
    let posts = []; // Array to store all post names

    const fetchHistory = async () => {
        const response = await fetch(
            `http://localhost:3000/vote/history/election?election_id=${electionId}`,
            {
                headers: {
                    Accept: "application/json",
                },
            },
        );
        if (response.ok) {
            historyData = await response.json();
            posts = Object.keys(historyData); // Store post names
            if (!initialRender) {
                updateGraphs(historyData);
            }
        } else {
            console.error("Failed to fetch history:", response.status);
        }
    };

    const processDataForChart = (postData) => {
        try {
            if (!postData) return { labels: [], datasets: [] };

            // Sort votes by timestamp
            const sortedVotes = postData.sort(
                (a, b) => parseFloat(a.timestamp) - parseFloat(b.timestamp),
            );

            const candidateVotes = {};
            const timeline = [];
            const datasets = [];

            sortedVotes.forEach((vote) => {
                const timestamp = new Date(parseFloat(vote.timestamp) * 1000);
                if (!timeline.includes(timestamp.getTime())) {
                    timeline.push(timestamp.getTime());
                }
                if (!candidateVotes[vote.candidate_id]) {
                    candidateVotes[vote.candidate_id] = [];
                }
                candidateVotes[vote.candidate_id].push(timestamp.getTime());
            });

            const sortedTimeline = timeline.sort((a, b) => a - b);
            const uniqueTimestamps = [...new Set(sortedTimeline)];
            const labels = uniqueTimestamps.map((ts) =>
                new Date(ts).toLocaleTimeString(),
            );

            for (const candidateId in candidateVotes) {
                const voteCounts = [];
                let count = 0;
                for (const ts of uniqueTimestamps) {
                    count += candidateVotes[candidateId].filter(
                        (voteTs) => voteTs <= ts,
                    ).length;
                    voteCounts.push(count);
                }
                datasets.push({
                    label: `Candidate ${candidateId}`,
                    data: voteCounts,
                    borderColor: getRandomColor(),
                    backgroundColor: (context) => {
                        const chart = context.chart;
                        const { ctx, chartArea } = chart;

                        if (!chartArea) {
                            return null;
                        }

                        const gradientBg = ctx.createLinearGradient(
                            0,
                            chartArea.bottom,
                            0,
                            chartArea.top,
                        );
                        gradientBg.addColorStop(
                            0,
                            datasets[datasets.length - 1].borderColor + "40",
                        );
                        gradientBg.addColorStop(1, "transparent");
                        return gradientBg;
                    },
                    fill: true,
                    tension: 0.3,
                });
            }

            return { labels, datasets };
        } catch (error) {
            console.error("Error in processDataForChart:", error);
            return { labels: [], datasets: [] };
        }
    };

    const updateGraphForPost = (post, data) => {
        const chartData = processDataForChart(data);
        const canvasId = `chart-${post.toLowerCase()}`;
        const canvasElement = canvasElements[canvasId]; // Access from stored elements

        if (canvasElement) {
            const ctx = canvasElement.getContext("2d");
            if (ctx) {
                if (chartInstances[post]) {
                    chartInstances[post].data = {
                        labels: chartData.labels,
                        datasets: chartData.datasets,
                    };
                    chartInstances[post].update();
                } else {
                    chartInstances[post] = new Chart(ctx, {
                        type: "line",
                        data: {
                            labels: chartData.labels,
                            datasets: chartData.datasets,
                        },
                        options: {
                            responsive: true,
                            maintainAspectRatio: false,
                            scales: {
                                x: {
                                    title: {
                                        display: true,
                                        text: "Time",
                                    },
                                },
                                y: {
                                    title: {
                                        display: true,
                                        text: "Total Votes",
                                    },
                                    beginAtZero: true,
                                },
                            },
                            plugins: {
                                title: {
                                    display: true,
                                    text: `Vote History for ${post}`,
                                },
                                legend: {
                                    display: true,
                                },
                            },
                        },
                    });
                }
            } else {
                console.error(
                    `Failed to get 2D context for canvas: ${canvasId}`,
                );
            }
        } else {
            console.log("Canvas element not found for", post);
        }
    };

    const updateGraphs = (data) => {
        for (const post of posts) {
            // Iterate over all posts
            if (canvasElements[`chart-${post.toLowerCase()}`]) {
                updateGraphForPost(post, data[post]);
            }
        }
    };

    const handleNewVote = (event) => {
        try {
            const message = JSON.parse(event.data);
            if (message.election_id === electionId) {
                for (const post in message.candidates) {
                    if (historyData[post]) {
                        historyData[post] = [
                            ...historyData[post],
                            {
                                candidate_id: message.candidates[post],
                                timestamp: Date.now() / 1000,
                            },
                        ];
                        updateGraphForPost(post, historyData[post]);
                    }
                }
            }
        } catch (error) {
            console.error("Error processing WebSocket message:", error);
        }
    };

    onMount(async () => {
        await fetchHistory();

        // Establish WebSocket connection only once
        if (!socket || socket.readyState === WebSocket.CLOSED) {
            socket = new WebSocket("ws://localhost:3000/vote/ws");

            socket.onopen = () => {
                console.log("WebSocket connected");
            };

            socket.onmessage = handleNewVote;

            socket.onclose = () => {
                console.log("WebSocket connection closed");
            };

            socket.onerror = (error) => {
                console.error("WebSocket error:", error);
            };
        }
        // Delay initial graph rendering
        setTimeout(() => {
            initialRender = false;
            updateGraphs(historyData);
        }, 0);

        return () => {
            if (socket && socket.readyState === WebSocket.OPEN) {
                socket.close();
            }
        };
    });

    const handlePostChange = (event) => {
        selectedPost = event.target.value;
        updateGraphForPost(selectedPost, historyData[selectedPost]);
    };

    function getRandomColor() {
        const letters = "0123456789ABCDEF";
        let color = "#";
        for (let i = 0; i < 6; i++) {
            color += letters[Math.floor(Math.random() * 16)];
        }
        return color;
    }
</script>

<svelte:head>
    <title>Vote History</title>
</svelte:head>

<Navbar />
<section class="m-12">
    <h1>Vote History for Election {electionId}</h1>

    <div>
        <label for="post-filter">Filter by Post:</label>
        <select
            id="post-filter"
            bind:value={selectedPost}
            on:change={handlePostChange}
        >
            {#each posts as post (post)}
                <option value={post}>{post}</option>
            {/each}
        </select>
    </div>

    <div style="width: 80%; margin: 0 auto;">
        {#each posts as post (post)}
            <div style="margin-bottom: 20px;">
                <h2>Vote History for {post}</h2>
                <canvas
                    id={`chart-${post.toLowerCase()}`}
                    bind:this={canvasElements[`chart-${post.toLowerCase()}`]}
                ></canvas>
            </div>
        {/each}
    </div>

    {#if Object.keys(historyData).length === 0}
        <p>Loading vote history...</p>
    {:else}
        <p>No history available for the selected post.</p>
    {/if}

    <style>
        h1 {
            text-align: center;
            margin-bottom: 20px;
        }

        div {
            margin-bottom: 15px;
        }

        label {
            margin-right: 10px;
            font-weight: bold;
        }

        select {
            padding: 8px;
            border: 1px solid #ccc;
            border-radius: 4px;
        }

        canvas {
            width: 100%;
            max-height: 400px;
        }

        p {
            font-style: italic;
            color: #777;
        }
        h2 {
            text-align: center;
        }
    </style>
</section>
