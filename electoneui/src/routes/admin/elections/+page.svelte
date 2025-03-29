<script>
    import { onMount } from "svelte";
    import {
        GetElections,
        DeleteElection,
        UpdateElection,
        CreateElection,
    } from "./api.js";

    let elections = [];
    let newElection = {
        title: "",
        startDate: "",
        endDate: "",
        votingRestriction: "", // Note the singular form
    };

    let editingElection = null;

    async function loadElections() {
        try {
            const data = await GetElections();
            elections = data || [];
        } catch (error) {
            console.error("Failed to load elections:", error);
        }
    }

    async function removeElection(electionID) {
        try {
            await DeleteElection(electionID);
            await loadElections();
        } catch (error) {
            console.error("Failed to delete election:", error);
        }
    }

    async function addElection() {
        try {
            await CreateElection(newElection);
            newElection = {
                title: "",
                startDate: "",
                endDate: "",
                votingRestriction: "",
            };
            await loadElections();
        } catch (error) {
            console.error("Failed to add election:", error);
        }
    }

    async function saveElection() {
        if (!editingElection) return;

        try {
            // Create a clean payload with only the fields the backend expects
            const payload = {
                title: editingElection.title,
                startDate: editingElection.startDate,
                endDate: editingElection.endDate,
                votingRestriction: editingElection.votingRestriction,
            };

            await UpdateElection(editingElection.electionID, payload);
            editingElection = null;
            await loadElections();
        } catch (error) {
            console.error("Failed to update election:", error);
        }
    }

    function startEditing(election) {
        editingElection = {
            electionID: election.electionID,
            title: election.title || "",
            startDate: election.startDate || "",
            endDate: election.endDate || "",
            votingRestriction: election.votingRestriction || "",
        };
    }

    onMount(loadElections);
</script>

<main>
    <h1>Elections</h1>
    <ul>
        {#each elections as election (election.electionID)}
            <li>
                {#if editingElection?.electionID === election.electionID}
                    <input
                        bind:value={editingElection.title}
                        placeholder="Title"
                    />
                    <input
                        bind:value={editingElection.startDate}
                        placeholder="Start Date"
                        type="date"
                    />
                    <input
                        bind:value={editingElection.endDate}
                        placeholder="End Date"
                        type="date"
                    />
                    <input
                        bind:value={editingElection.votingRestriction}
                        placeholder="Restrictions"
                    />
                    <button on:click={saveElection}>Save</button>
                {:else}
                    {election.electionID} - {election.title} - {election.startDate} - {election.endDate}
                    - {election.votingRestriction}
                    <button on:click={() => startEditing(election)}>Edit</button
                    >
                    <button on:click={() => removeElection(election.electionID)}
                        >Delete</button
                    >
                {/if}
            </li>
        {/each}
    </ul>

    <h2>Add Election</h2>
    <input bind:value={newElection.title} placeholder="Title" />
    <input
        bind:value={newElection.startDate}
        placeholder="Start Date"
        type="date"
    />
    <input
        bind:value={newElection.endDate}
        placeholder="End Date"
        type="date"
    />
    <input
        bind:value={newElection.votingRestriction}
        placeholder="Restrictions"
    />
    <button on:click={addElection}>Add</button>
</main>
