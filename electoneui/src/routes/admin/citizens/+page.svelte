<script>
    import { onMount } from "svelte";
    import {
        GetCitizens,
        DeleteCitizen,
        UpdateCitizen,
        CreateCitizen,
    } from "./api.js";

    let citizens = [];
    let searchQuery = "";
    let newCitizen = {
        citizenID: "",
        fullName: "",
        dateOfBirth: "",
        placeOfResidence: "",
    };
    let editingCitizen = null;

    async function loadCitizens() {
        citizens = await GetCitizens();
    }

    async function removeCitizen(id) {
        await DeleteCitizen(id);
        loadCitizens();
    }

    async function addCitizen() {
        await CreateCitizen(newCitizen);
        newCitizen = {
            citizenID: "",
            fullName: "",
            dateOfBirth: "",
            placeOfResidence: "",
        };
        loadCitizens();
    }

    async function saveCitizen() {
        if (editingCitizen) {
            await UpdateCitizen(editingCitizen.citizenID, editingCitizen);
            editingCitizen = null;
            loadCitizens();
        }
    }

    onMount(loadCitizens);
</script>

<main class="mx-auto p-10 bg-black text-gray-300">
    <div class="flex justify-between items-center mb-4">
        <input
            type="text"
            bind:value={searchQuery}
            placeholder="Search by Citizen ID..."
            class="input input-bordered bg-[#0a0a0a] text-white"
        />
        <button class="btn btn-primary" on:click={() => showModal.showModal()}
            >+ New Record</button
        >
    </div>

    <div class="overflow-x-auto flex justify-center">
        <table
            class="table w-full border border-[#070707] bg-[#0a0a0a] text-white"
        >
            <thead>
                <tr class="bg-[#070707]">
                    <th class="text-center">Citizen ID</th>
                    <th class="text-center">Full Name</th>
                    <th class="text-center">Date of Birth</th>
                    <th class="text-center">Place of Residence</th>
                    <th class="text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each citizens.filter((c) => c.citizenID
                        .toLowerCase()
                        .includes(searchQuery.toLowerCase())) as citizen}
                    <tr class="hover">
                        {#if editingCitizen?.citizenID === citizen.citizenID}
                            <td
                                ><input
                                    bind:value={editingCitizen.citizenID}
                                    disabled
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCitizen.fullName}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCitizen.dateOfBirth}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td
                                ><input
                                    bind:value={editingCitizen.placeOfResidence}
                                    class="input input-bordered bg-[#0a0a0a] text-white"
                                /></td
                            >
                            <td>
                                <button
                                    class="btn btn-primary"
                                    on:click={saveCitizen}>Save</button
                                >
                            </td>
                        {:else}
                            <td class="text-center border-r border-white"
                                >{citizen.citizenID}</td
                            >
                            <td class="text-center border-r border-white"
                                >{citizen.fullName}</td
                            >
                            <td class="text-center border-r border-white"
                                >{citizen.dateOfBirth}</td
                            >
                            <td class="text-center"
                                >{citizen.placeOfResidence}</td
                            >
                            <td>
                                <button
                                    class="btn btn-secondary text-black"
                                    on:click={() =>
                                        (editingCitizen = { ...citizen })}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-error"
                                    on:click={() =>
                                        removeCitizen(citizen.citizenID)}
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
            <h2 class="text-xl mb-4">Add New Citizen</h2>
            <input
                bind:value={newCitizen.citizenID}
                placeholder="Citizen ID"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newCitizen.fullName}
                placeholder="Full Name"
                class="input input-bordered w-full bg-black text-white mb-2"
            />
            <input
                bind:value={newCitizen.dateOfBirth}
                placeholder="Date of Birth"
                class="input input-bordered w-full bg-black text-white mb-2"
                type="date"
            />
            <input
                bind:value={newCitizen.placeOfResidence}
                placeholder="Place of Residence"
                class="input input-bordered w-full bg-black text-white mb-4"
            />

            <div class="flex justify-end space-x-2">
                <form method="dialog">
                    <button class="btn btn-accent text-black">Cancel</button>
                </form>
                <button class="btn btn-primary" on:click={addCitizen}
                    >Create</button
                >
            </div>
        </div>
    </dialog>
</main>
