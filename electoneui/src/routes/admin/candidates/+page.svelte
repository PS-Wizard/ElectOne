<script>
    import { onMount } from "svelte";
    import {
        GetCandidates,
        DeleteCandidate,
        UpdateCandidate,
        CreateCandidate,
    } from "./api.js";

    let candidates = [];
    let searchQuery = "";
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
        closeModal(); // Close modal after adding candidate
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

    function startEditing(candidate) {
        editingCandidate = { ...candidate };
    }

    function openModal() {
        const modal = document.getElementById("showModal");
        modal.showModal();
    }

    function closeModal() {
        const modal = document.getElementById("showModal");
        modal.close();
    }

    onMount(loadCandidates);
</script>

<main class="mx-auto p-10 bg-black text-gray-300">
    <div class="flex justify-between items-center mb-4">
        <input
            type="text"
            bind:value={searchQuery}
            placeholder="Search by Citizen ID..."
            class="input input-bordered bg-[#0a0a0a] text-white"
        />
        <button class="btn btn-primary" on:click={openModal}>
            + New Record
        </button>
    </div>

    <div class="overflow-x-auto flex justify-center">
        <table
            class="table w-full border border-[#070707] bg-[#0a0a0a] text-white"
        >
            <thead>
                <tr class="bg-[#070707]">
                    <th class="text-center">Candidate ID</th>
                    <th class="text-center">Citizen ID</th>
                    <th class="text-center">Post</th>
                    <th class="text-center">Election ID</th>
                    <th class="text-center">Group</th>
                    <th class="text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each candidates.filter((candidate) => candidate.citizenID
                        .toLowerCase()
                        .includes(searchQuery.toLowerCase())) as candidate}
                    <tr class="hover">
                        {#if editingCandidate?.candidateID === candidate.candidateID}
                            <td
                                ><input
                                    bind:value={editingCandidate.candidateID}
                                    disabled
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCandidate.citizenID}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCandidate.post}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCandidate.electionID}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCandidate.group}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td>
                                <button
                                    class="btn btn-primary"
                                    on:click={saveCandidate}>Save</button
                                >
                            </td>
                        {:else}
                            <td class="text-center border-r border-white"
                                >{candidate.candidateID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{candidate.citizenID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{candidate.post}</td
                            >
                            <td class="text-center border-r border-white"
                                >{candidate.electionID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{candidate.group}</td
                            >
                            <td>
                                <button
                                    class="btn btn-secondary text-black"
                                    on:click={() => startEditing(candidate)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-error"
                                    on:click={() =>
                                        removeCandidate(candidate.candidateID)}
                                    >Delete</button
                                >
                            </td>
                        {/if}
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>

    <!-- Modal for New Candidate -->
    <dialog id="showModal" class="modal">
        <div class="modal-box bg-[#0a0a0a] text-white">
            <h2 class="text-xl mb-4">Add New Candidate</h2>
            <input
                bind:value={newCandidate.citizenID}
                placeholder="Citizen ID"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newCandidate.post}
                placeholder="Post"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newCandidate.electionID}
                placeholder="Election ID"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newCandidate.group}
                placeholder="Group"
                class="input input-bordered w-full bg-black text-white mb-4"
            />

            <div class="flex justify-end space-x-2">
                <form method="dialog">
                    <button class="btn btn-secondary" on:click={closeModal}
                        >Cancel</button
                    >
                </form>
                <button class="btn btn-primary" on:click={addCandidate}
                    >Create</button
                >
            </div>
        </div>
    </dialog>
</main>
