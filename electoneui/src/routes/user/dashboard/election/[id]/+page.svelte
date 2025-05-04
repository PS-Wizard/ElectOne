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
        const token = localStorage.getItem("token");
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
            if (!groupedCandidates[candidate.post]) {
                groupedCandidates[candidate.post] = [];
            }
            groupedCandidates[candidate.post].push(candidate);
        });
    }

    // Function to handle voting submission
    async function submitVote() {
        const token = localStorage.getItem("token");

        // Make sure at least one candidate per post is selected
        if (
            Object.keys(selectedCandidates).length === 0 ||
            Object.keys(selectedCandidates).length !==
                Object.keys(groupedCandidates).length
        ) {
            alert("Please select a candidate for each post.");
            return;
        }

        const votePayload = {
            election_id: parseInt(electionId),
            candidates: selectedCandidates, // candidates as a map
        };

        try {
            const res = await fetch("http://localhost:3000/vote", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(votePayload),
            });

            if (!res.ok) throw new Error("Failed to submit vote");

            alert("Vote submitted successfully!");
        } catch (err) {
            alert(err.message);
        }
    }
</script>

<Navbar />
<section class="m-12">
    <h1 class="text-2xl font-bold mb-4">Election #{electionId}</h1>

    {#if loading}
        <p>Loading candidates...</p>
    {:else if error}
        <p class="text-red-500">Error: {error}</p>
    {:else if Object.keys(groupedCandidates).length === 0}
        <p>No candidates found for this election.</p>
    {:else}
        <div>
            {#each Object.entries(groupedCandidates) as [post, postCandidates]}
                <div class="mb-6">
                    <h2 class="text-xl font-semibold">{post}</h2>
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                        {#each postCandidates as candidate (candidate.candidate_id)}
                            <label
                                class="card bg-base-200 shadow-lg p-4 space-y-2 cursor-pointer"
                                on:click={() =>
                                    (selectedCandidates[post] =
                                        candidate.candidate_id)}
                            >
                                <input
                                    type="radio"
                                    name={post}
                                    value={candidate.candidate_id}
                                    bind:group={selectedCandidates[post]}
                                    class="invisible"
                                />
                                <div class="flex items-center space-x-4">
                                    <img
                                        src={candidate.profile_path}
                                        alt={candidate.name}
                                        class="w-16 h-16 rounded-full"
                                    />
                                    <div class="text-left">
                                        <h3 class="font-semibold">
                                            {candidate.name}
                                        </h3>
                                        <p class="text-sm text-gray-500">
                                            {candidate.party}
                                        </p>
                                        <p class="text-sm">{candidate.bio}</p>
                                    </div>
                                </div>
                            </label>
                        {/each}
                    </div>
                </div>
            {/each}

            <button class="btn btn-primary mt-4" on:click={submitVote}>
                Submit Vote
            </button>
        </div>
    {/if}
</section>
