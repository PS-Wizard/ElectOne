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

    $: percentage =
        totalVotesInPost > 0
            ? ((candidate.votes / totalVotesInPost) * 100).toFixed(1)
            : 0;

    onMount(() => {
        if (canvasElement) {
            const ctx = canvasElement.getContext("2d");
            chartInstance = new Chart(ctx, {
                type: "pie",
                data: {
                    labels: ["This Candidate", "Others in Post"],
                    datasets: [
                        {
                            data: [
                                candidate.votes,
                                totalVotesInPost - candidate.votes,
                            ],
                            backgroundColor: [
                                "rgba(75, 192, 192, 0.7)", // Candidate's color
                                "rgba(201, 203, 207, 0.7)", // Others' color
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
                                    const total = context.dataset.data.reduce(
                                        (a, b) => a + b,
                                        0,
                                    );
                                    const currentPercentage = Math.round(
                                        (value / total) * 100,
                                    );
                                    return `${label}: ${value} (${currentPercentage}%)`;
                                },
                            },
                        },
                    },
                },
            });
        }
        return () => {
            if (chartInstance) {
                chartInstance.destroy();
            }
        };
    });
</script>

<div class="border shadow-sm border-black justify-center items-center flex flex-col p-4 rounded-lg cursor-pointer transition-transform duration-200 hover:-translate-y-[12px]">
    <img
        src={candidate.candidate_photo}
        alt={candidate.candidate_name}
        class="candidate-image border w-48 h-48 rounded-full object-cover "
    />
    <h3 class="font-medium text-xl">{candidate.candidate_name}</h3>
    <p class="party">{candidate.candidate_party}</p>
    <div class="vote-percentage">{percentage}%</div>
    <div class="chart-container">
        <canvas bind:this={canvasElement}></canvas>
    </div>
</div>

<style>
</style>
