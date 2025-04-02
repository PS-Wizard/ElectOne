<script>
    import { onMount } from "svelte";
    import {
        GetElections,
        DeleteElection,
        UpdateElection,
        CreateElection,
    } from "./api.js";

    let elections = [];
    let searchQuery = "";
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
            closeModal(); // Close modal after adding election
        } catch (error) {
            console.error("Failed to add election:", error);
        }
    }

    async function saveElection() {
        if (!editingElection) return;

        try {
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

    function openModal() {
        const modal = document.getElementById("showModal");
        modal.showModal();
    }

    function closeModal() {
        const modal = document.getElementById("showModal");
        modal.close();
    }

    onMount(loadElections);
</script>

<main class="mx-auto p-10 bg-black text-gray-300">
    <div class="flex justify-between items-center mb-4">
        <input
            type="text"
            bind:value={searchQuery}
            placeholder="Search by Title..."
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
                    <th class="text-center">Election ID</th>
                    <th class="text-center">Title</th>
                    <th class="text-center">Start Date</th>
                    <th class="text-center">End Date</th>
                    <th class="text-center">Voting Restriction</th>
                    <th class="text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each elections.filter((e) => e.title
                        .toLowerCase()
                        .includes(searchQuery.toLowerCase())) as election}
                    <tr class="hover">
                        {#if editingElection?.electionID === election.electionID}
                            <td
                                ><input
                                    bind:value={editingElection.electionID}
                                    disabled
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingElection.title}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingElection.startDate}
                                    type="date"
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingElection.endDate}
                                    type="date"
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={
                                        editingElection.votingRestriction
                                    }
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td>
                                <button
                                    class="btn btn-primary"
                                    on:click={saveElection}>Save</button
                                >
                            </td>
                        {:else}
                            <td class="text-center border-r border-white"
                                >{election.electionID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{election.title}</td
                            >
                            <td class="text-center border-r border-white"
                                >{election.startDate}</td
                            >
                            <td class="text-center border-r border-white"
                                >{election.endDate}</td
                            >
                            <td class="text-center border-r border-white"
                                >{election.votingRestriction}</td
                            >
                            <td>
                                <button
                                    class="btn btn-secondary text-black"
                                    on:click={() => startEditing(election)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-error"
                                    on:click={() =>
                                        removeElection(election.electionID)}
                                    >Delete</button
                                >
                            </td>
                        {/if}
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>

    <dialog id="showModal" class="modal">
        <div class="modal-box bg-[#0a0a0a] text-white">
            <h2 class="text-xl mb-4">Add New Election</h2>
            <input
                bind:value={newElection.title}
                placeholder="Title"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newElection.startDate}
                type="date"
                placeholder="Start Date"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newElection.endDate}
                type="date"
                placeholder="End Date"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newElection.votingRestriction}
                placeholder="Restrictions"
                class="input input-bordered w-full bg-black text-white mb-4"
            />

            <div class="flex justify-end space-x-2">
                <form method="dialog">
                    <button class="btn btn-secondary" on:click={closeModal}
                        >Cancel</button
                    >
                </form>
                <button class="btn btn-primary" on:click={addElection}
                    >Create</button
                >
            </div>
        </div>
    </dialog>
</main>
