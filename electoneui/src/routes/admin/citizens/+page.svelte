<script>
    import { onMount } from "svelte";
    import {
        GetCitizens,
        DeleteCitizen,
        UpdateCitizen,
        CreateCitizen,
    } from "./api.js";

    let citizens = [];
    let newCitizen = {
        citizenID: "",
        fullName: "",
        dateOfBirth: "",
        placeOfResidence: "",
    };
    let editingCitizen = null;

    async function loadCitizens() {
        citizens = await GetCitizens();
        console.log(citizens);
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

<main>
    <h1>Citizens</h1>
    <ul>
        {#each citizens as citizen}
            <li>
                {#if editingCitizen?.citizenID === citizen.citizenID}
                    <input
                        bind:value={editingCitizen.fullName}
                        placeholder="Full Name"
                    />
                    <input
                        bind:value={editingCitizen.dateOfBirth}
                        placeholder="DOB"
                    />
                    <input
                        bind:value={editingCitizen.placeOfResidence}
                        placeholder="Residence"
                    />
                    <button on:click={saveCitizen}>Save</button>
                {:else}
                    {citizen.fullName} - {citizen.dateOfBirth} - {citizen.placeOfResidence}
                    <button on:click={() => (editingCitizen = { ...citizen })}
                        >Edit</button
                    >
                    <button on:click={() => removeCitizen(citizen.citizenID)}
                        >Delete</button
                    >
                {/if}
            </li>
        {/each}
    </ul>
    <h2>Add Citizen</h2>
    <input bind:value={newCitizen.citizenID} placeholder="ID" />
    <input bind:value={newCitizen.fullName} placeholder="Full Name" />
    <input bind:value={newCitizen.dateOfBirth} placeholder="DOB" />
    <input bind:value={newCitizen.placeOfResidence} placeholder="Residence" />
    <button on:click={addCitizen}>Add</button>
</main>
