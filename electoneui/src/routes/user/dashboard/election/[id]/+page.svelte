<script>
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import Navbar from "../../../../../components/Navbar.svelte";

    let electionId;
    let candidates = [];
    let groupedCandidates = {};
    let selectedCandidates = {};
    let loading = true;
    let error = null;

    $: electionId = $page.params.id;

    onMount(async () => {
        const token = localStorage.getItem("user_token");
        if (!token) {
            error = "You're not logged in.";
            loading = false;
            return;
        }

        try {
            const res = await fetch(
                "http://localhost:3000/candidate?limit=999&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch candidates");

            const data = await res.json();
            console.log(data);
            candidates = data.filter((c) => c.election_id == electionId);
            groupCandidatesByPost(candidates);
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    });

    function groupCandidatesByPost(candidates) {
        candidates.forEach((candidate) => {
            if (!groupedCandidates[candidate.candidate_post]) {
                groupedCandidates[candidate.candidate_post] = [];
            }
            groupedCandidates[candidate.candidate_post].push(candidate);
        });
    }

    // Function to handle voting submission
    async function submitVote() {
        const token = localStorage.getItem("user_token");

        if (
            Object.keys(selectedCandidates).length === 0 ||
            Object.keys(selectedCandidates).length !==
                Object.keys(groupedCandidates).length
        ) {
            alert("Please select a candidate for each post.");
            return;
        }

        try {
            for (const post in selectedCandidates) {
                const candidateID = selectedCandidates[post];

                const res = await fetch("http://localhost:3000/vote", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${token}`,
                    },
                    body: JSON.stringify({
                        election_id: parseInt(electionId),
                        candidate_id: candidateID,
                    }),
                });

                const data = await res.json();

                if (!res.ok) {
                    throw new Error(data.error || "Failed to submit vote.");
                }

                alert(
                    "Vote for candidate " +
                        candidateID +
                        " submitted successfully!",
                );
            }
        } catch (err) {
            alert(err.message);
        }
    }
</script>

<Navbar />
<section class="m-4 md:m-12">
    <h1 class="text-3xl sm:text-6xl font-bold text-center mb-4">
        Pick Your Votes:
    </h1>

    {#if loading}
        <p>Loading candidates...</p>
    {:else if error}
        <p class="text-red-500">Error: {error}</p>
    {:else if Object.keys(groupedCandidates).length === 0}
        <p>No candidates found for this election.</p>
    {:else}
        <div class="flex flex-col items-center">
            {#each Object.entries(groupedCandidates) as [post, postCandidates]}
                <div class="m-6">
                    <h2 class="text-lg sm:text-xl font-semibold uppercase mb-4">
                        {post}
                    </h2>
                    <div
                        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"
                    >
                        {#each postCandidates as candidate (candidate.candidate_id)}
                            <label
                                class="card bg-base-200 border p-4 cursor-pointer hover:shadow-lg transition-all transform hover:scale-105"
                                on:click={() =>
                                    (selectedCandidates[post] =
                                        candidate.candidate_id)}
                            >
                                <input
                                    type="radio"
                                    name={post}
                                    value={candidate.candidate_id}
                                    bind:group={selectedCandidates[post]}
                                    class="hidden"
                                />

                                <div
                                    class="flex flex-col sm:flex-row items-center sm:space-x-4"
                                >
                                    <!-- Profile on the left (sm:w-1/3) -->
                                    <div
                                        class="flex-shrink-0 sm:w-1/3 mb-4 sm:mb-0"
                                    >
                                        <img
                                            src={`http://localhost:3000${candidate.candidate_photo}`}
                                            alt={candidate.candidate_name}
                                            class="w-full h-auto rounded-full border-2 border-primary object-cover"
                                        />
                                    </div>

                                    <!-- Candidate info on the right (sm:w-2/3) -->
                                    <div class="sm:w-2/3">
                                        <h3
                                            class="font-bold text-lg sm:text-xl"
                                        >
                                            {candidate.candidate_name}
                                        </h3>
                                        <p class="text-sm text-gray-500">
                                            {candidate.candidate_party}
                                        </p>
                                        <p class="text-sm text-gray-600 mt-2">
                                            {candidate.candidate_bio}
                                        </p>
                                    </div>
                                </div>
                            </label>
                        {/each}
                    </div>
                </div>
            {/each}

            <button
                class="btn btn-primary mt-4 min-w-1/2 sm:w-auto"
                on:click={submitVote}
            >
                Submit Vote
            </button>
        </div>
    {/if}
</section>
