<script>
    import { onMount, onDestroy } from "svelte"; 
    import {
        Chart,
        PieController,
        ArcElement,
        Tooltip,
        Legend,
    } from "chart.js";

    Chart.register(PieController, ArcElement, Tooltip, Legend);

    export let candidate;
    export let totalVotesInPost;

    let canvasElement;
    let chartInstance = null;

    $: if (
        candidate &&
        typeof totalVotesInPost !== "undefined" &&
        canvasElement
    ) {
        console.log(
            `PieChart Update Triggered for: ${candidate.candidate_name}, Votes: ${candidate.votes}, Total: ${totalVotesInPost}`,
        );
        updateChart();
    }

    $: percentage =
        totalVotesInPost > 0
            ? ((candidate.votes / totalVotesInPost) * 100).toFixed(1)
            : 0;

    function updateChart() {
        if (!canvasElement) {
            console.warn("Canvas element not yet available for pie chart.");
            return;
        }

        if (chartInstance) {
            chartInstance.destroy();
            chartInstance = null;
        }

        const candidateVotes = Number(candidate.votes) || 0;
        const totalVotes = Number(totalVotesInPost) || 0;
        const otherVotes = Math.max(0, totalVotes - candidateVotes);

        const ctx = canvasElement.getContext("2d");
        chartInstance = new Chart(ctx, {
            type: "pie",
            data: {
                labels: ["This Candidate", "Others in Post"],
                datasets: [
                    {
                        data: [
                            candidateVotes,
                            otherVotes, 
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
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: "bottom",
                    },
                    tooltip: {
                        callbacks: {
                            label: function (context) {
                                const label = context.label || "";
                                const value = context.raw || 0;
                                const currentDatasetTotal =
                                    context.dataset.data.reduce(
                                        (a, b) => a + b,
                                        0,
                                    );
                                const currentPercentage =
                                    currentDatasetTotal > 0
                                        ? Math.round(
                                              (value / currentDatasetTotal) *
                                                  100,
                                          )
                                        : 0;
                                return `${label}: ${value} (${currentPercentage}%)`;
                            },
                        },
                    },
                },
            },
        });
    }

    onMount(() => {
        if (
            canvasElement &&
            candidate &&
            typeof totalVotesInPost !== "undefined"
        ) {
            console.log(`PieChart onMount for: ${candidate.candidate_name}`);
            updateChart();
        }

        return () => {
            if (chartInstance) {
                console.log(
                    `PieChart onDestroy for: ${candidate.candidate_name}`,
                );
                chartInstance.destroy();
                chartInstance = null;
            }
        };
    });
</script>

<div
    class="border shadow-sm border-black justify-center items-center flex flex-col p-4 rounded-lg cursor-pointer transition-transform duration-200 hover:-translate-y-[12px]"
>
    <img
        src={candidate.candidate_photo}
        alt={candidate.candidate_name}
        class="candidate-image border w-32 h-32 md:w-48 md:h-48 rounded-full object-cover mb-3"
    />
    <h3 class="font-medium text-lg md:text-xl text-center">
        {candidate.candidate_name}
    </h3>
    <p class="party text-sm text-gray-600">{candidate.candidate_party}</p>
    <div class="vote-percentage font-semibold text-blue-600 my-1">
        {percentage}%
    </div>
    <div class="chart-container w-full h-40 md:h-48 mt-2">
        <canvas bind:this={canvasElement}></canvas>
    </div>
</div>

<style>
    .chart-container {
        position: relative; 
    }
    .chart-container canvas {
        max-width: 100%;
        max-height: 100%;
    }

    .candidate-image {
        flex-shrink: 0; 
    }
</style>
