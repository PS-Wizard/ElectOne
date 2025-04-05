<script>
    import { onMount, onDestroy } from "svelte";
    import { writable, derived } from "svelte/store";
    import { browser } from "$app/environment";

    let elections = writable([]);
    let votes = writable({});
    let selectedCandidates = writable({}); // Track selected candidates per post
    let candidatesByPost = writable({});
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
        console.log("$candidatesByPost:", grouped); // Add this line
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

    let socket;

    const candidateVotes = derived(
        [votes, candidatesByPost],
        ([$votes, $candidatesByPost]) => {
            const votesMap = {};
            // Flatten all candidates and map their votes
            Object.values($candidatesByPost)
                .flat()
                .forEach((candidate) => {
                    votesMap[candidate.candidateID] =
                        $votes[candidate.candidateID] || 0;
                });
            return votesMap;
        },
    );

    onMount(() => {
        fetchElections();

        // Connect to WebSocket
        socket = new WebSocket("wss://localhost:3000/api/voteStream");

        // Modify your WebSocket handler to force a new object reference
        socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log($votes);
            console.log(data);
            votes.update((v) => {
                return { ...v, [data.candidateID]: data.votes };
            });
        };

        socket.onerror = (err) => {
            console.error("WebSocket error:", err);
        };

        socket.onclose = () => {
            console.log("WebSocket connection closed");
        };
    });

    onDestroy(() => {
        if (socket) socket.close();
        if (chart) chart.destroy();
    });
</script>

<h1 class="text-3xl font-bold text-center mb-6">Users Page Dashboard</h1>
<!-- Grouped Candidates with Select Options -->
<div class="space-y-6">
    {#each $elections as election}
        <div class="card shadow-lg p-6 rounded-lg">
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
                    <div class="flex flex-col gap-4">
                        <h4 class="text-lg font-semibold">{post}</h4>
                        <div class="flex flex-wrap gap-4">
                            {#each candidates as candidate (candidate.candidateID)}
                                <div
                                    class="relative rounded-lg transition duration-300 transform hover:scale-105"
                                >
                                    <input
                                        type="radio"
                                        id={candidate.candidateID}
                                        name={post}
                                        class="peer hidden"
                                        bind:group={$selectedCandidates[post]}
                                        value={candidate.candidateID}
                                    />
                                    <label
                                        for={candidate.candidateID}
                                        class="card bg-base-100 w-80 shadow-md cursor-pointer group peer-checked:bg-white peer-checked:text-black peer-checked:border-black transition duration-300 transform hover:scale-105"
                                    >
                                        <figure>
                                            <img
                                                src="/profileplaceholder.jpg"
                                                alt={candidate.fullName}
                                                class="object-cover h-40 w-full rounded-t-lg"
                                            />
                                        </figure>
                                        <div class="card-body">
                                            <h2 class="card-title text-center">
                                                {candidate.fullName}
                                            </h2>
                                            <p
                                                class="text-center text-sm opacity-70"
                                            >
                                                Group: <span
                                                    class="badge badge-accent"
                                                >
                                                    {candidate.group}</span
                                                >
                                            </p>
                                            <p
                                                class="text-center text-lg font-bold"
                                            >
                                                Votes: <span
                                                    class="text-primary"
                                                >
                                                    {$votes[
                                                        String(
                                                            candidate.candidateID,
                                                        )
                                                    ] || 0}
                                                </span>
                                            </p>
                                        </div>
                                    </label>
                                </div>
                            {/each}
                        </div>
                    </div>
                {/each}
            </div>
            <div class="mt-6 flex justify-center">
                <button
                    class="btn btn-primary w-full sm:w-auto"
                    on:click={() => castVotes(election.electionID)}
                >
                    Cast Vote
                </button>
            </div>
        </div>
    {/each}
</div>
