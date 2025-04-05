<script>
    import { onMount, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import { browser } from "$app/environment";

    let elections = writable([]);
    let candidatesByPost = writable({});
    let selectedCandidates = writable({}); // Track selected candidates per post
    let chart; // For storing chart instance

    async function fetchElections() {
        const res = await fetch(
            "https://localhost:3000/api/secure/electionsPaginated/0",
            { credentials: "include" },
        );
        const data = await res.json();
        elections.set(data);
        fetchCandidates();
    }

    async function fetchCandidates() {
        const res = await fetch(
            "https://localhost:3000/api/secure/candidatesPaginated/0",
            { credentials: "include" },
        );
        const data = await res.json();
        const grouped = {};

        for (const candidate of data) {
            if (!grouped[candidate.post]) {
                grouped[candidate.post] = [];
            }
            const citizenRes = await fetch(
                `https://localhost:3000/api/secure/citizens/${candidate.citizenID}`,
                { credentials: "include" },
            );
            const citizenData = await citizenRes.json();
            grouped[candidate.post].push({ ...candidate, ...citizenData });
        }
        candidatesByPost.set(grouped);
    }

    async function castVotes(electionID) {
        const selected = $selectedCandidates;
        let allVotesSuccessful = true;

        for (const post in selected) {
            const candidateID = selected[post];
            const response = await fetch(
                "https://localhost:3000/api/castVote",
                {
                    method: "POST",
                    credentials: "include",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        candidateID: candidateID,
                        electionID: electionID,
                    }),
                },
            );
            const data = await response.json();
            if (!data.message) {
                allVotesSuccessful = false;
            }
        }

        if (allVotesSuccessful) {
            alert("Your Vote Has Been Captured");
        } else {
            alert("Some votes failed to be cast. Please try again.");
        }
    }

    function renderGraph() {
        if (browser) {
            const ctx = document
                .getElementById("multiLineChart")
                .getContext("2d");

            const data = {
                labels: ["A", "B", "C", "D", "E"],
                datasets: [
                    {
                        label: "Group 1",
                        data: [12, 19, 3, 5, 2],
                        borderColor: "#4F46E5",
                        fill: false,
                        tension: 0.1,
                    },
                    {
                        label: "Group 2",
                        data: [5, 9, 7, 8, 3],
                        borderColor: "#F59E0B",
                        fill: false,
                        tension: 0.1,
                    },
                    {
                        label: "Group 3",
                        data: [10, 15, 8, 6, 9],
                        borderColor: "#10B981",
                        fill: false,
                        tension: 0.1,
                    },
                ],
            };

            const options = {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    y: { beginAtZero: true },
                },
            };

            if (chart) {
                chart.destroy();
            }

            chart = new Chart(ctx, {
                type: "line",
                data: data,
                options: options,
            });
        }
    }

    onMount(() => {
        fetchElections();
        setTimeout(renderGraph, 500); // Delay renderGraph to make sure data is ready
    });

    onDestroy(() => {
        if (chart) chart.destroy();
    });
</script>

<h1 class="text-3xl font-bold text-center mb-6">Users Page Dashboard</h1>

<!-- Multi-line graph above -->
<div class="w-full h-64 mb-8">
    <canvas id="multiLineChart"></canvas>
</div>

<!-- Grouped Candidates with Select Options -->
<div class="space-y-6">
    {#each $elections as election}
        <div class="card bg-base-200 shadow-lg p-6 rounded-lg">
            <h2 class="text-2xl font-semibold">{election.title}</h2>
            <p class="text-sm text-gray-500">
                Start Date: {election.startDate} | End Date: {election.endDate}
            </p>
            <p class="text-sm text-gray-500">
                Restriction: {election.votingRestriction}
            </p>
            <h3 class="mt-6 text-lg font-semibold">Candidates:</h3>
            <div class="space-y-4">
                {#each Object.entries($candidatesByPost) as [post, candidates]}
                    <div>
                        <h4 class="text-lg font-semibold">{post}</h4>
                        <div class="space-y-2">
                            {#each candidates as candidate (candidate.candidateID)}
                                <div class="flex items-center space-x-4">
                                    <input
                                        type="radio"
                                        name={post}
                                        id={candidate.candidateID}
                                        class="radio"
                                        bind:group={$selectedCandidates[post]}
                                        value={candidate.candidateID}
                                    />
                                    <label for={candidate.candidateID}>
                                        {candidate.fullName} - {candidate.group}
                                    </label>
                                </div>
                            {/each}
                        </div>
                    </div>
                {/each}
            </div>
            <div class="mt-4">
                <!-- Passing the electionID to the castVotes function -->
                <button
                    class="btn btn-primary"
                    on:click={() => castVotes(election.electionID)}
                >
                    Cast Vote
                </button>
            </div>
        </div>
    {/each}
</div>
