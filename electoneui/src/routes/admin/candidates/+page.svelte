<script>
    import { onMount } from "svelte";
    import {
        GetCandidates,
        DeleteCandidate,
        UpdateCandidate,
        CreateCandidate,
    } from "./api.js";

    let candidates = [];
    let newCandidate = {
        citizenID: "",
        post: "",
        electionID: "",
        group: "",
    };
    let editingCandidate = null;

    async function loadCandidates() {
        candidates = (await GetCandidates()) || [];
    }

    async function removeCandidate(id) {
        await DeleteCandidate(id);
        loadCandidates();
    }

    async function addCandidate() {
        await CreateCandidate(newCandidate);
        newCandidate = { citizenID: "", post: "", electionID: "", group: "" };
        loadCandidates();
    }

    async function saveCandidate() {
        if (editingCandidate?.candidateID) {
            await UpdateCandidate(editingCandidate.candidateID, {
                citizenID: editingCandidate.citizenID,
                post: editingCandidate.post,
                electionID: editingCandidate.electionID,
                group: editingCandidate.group,
            });
            editingCandidate = null;
            loadCandidates();
        }
    }

    onMount(loadCandidates);
</script>

<main>
    <h1>Candidates</h1>
    <ul>
        {#each candidates as candidate}
            <li>
                {#if editingCandidate?.candidateID === candidate.candidateID}
                    <input
                        bind:value={editingCandidate.citizenID}
                        placeholder="Citizen ID"
                    />
                    <input
                        bind:value={editingCandidate.post}
                        placeholder="Post"
                    />
                    <input
                        bind:value={editingCandidate.electionID}
                        placeholder="Election ID"
                    />
                    <input
                        bind:value={editingCandidate.group}
                        placeholder="Group"
                    />
                    <button on:click={saveCandidate}>Save</button>
                {:else}
                    {candidate.candidateID} - {candidate.citizenID} - {candidate.post} - {candidate.electionID}
                    - {candidate.group}
                    <button
                        on:click={() => (editingCandidate = { ...candidate })}
                        >Edit</button
                    >
                    <button
                        on:click={() => removeCandidate(candidate.candidateID)}
                        >Delete</button
                    >
                {/if}
            </li>
        {/each}
    </ul>

    <h2>Add Candidate</h2>
    <input bind:value={newCandidate.citizenID} placeholder="Citizen ID" />
    <input bind:value={newCandidate.post} placeholder="Post" />
    <input bind:value={newCandidate.electionID} placeholder="Election ID" />
    <input bind:value={newCandidate.group} placeholder="Group" />
    <button on:click={addCandidate}>Add</button>
</main>
