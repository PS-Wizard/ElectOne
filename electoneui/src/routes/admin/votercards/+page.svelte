<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let voterCards = [];
    let filteredVoterCards = [];
    let loading = true;
    let error = "";
    let search = "";

    let newVoterCard = {
        voter_card_id: "",
        citizenship_id: "",
        location: "",
    };

    let editingVoterCard = null;
    let newVoterCardModal;
    let editVoterCardModal;

    async function fetchVoterCards() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/voter_card?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch voter cards");
            voterCards = await res.json();
            filteredVoterCards = voterCards; // Initially, no filtering
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    function filterVoterCards() {
        filteredVoterCards = voterCards.filter((voterCard) => {
            const searchLower = search.toLowerCase();
            return (
                voterCard.voter_card_id.toLowerCase().includes(searchLower) ||
                voterCard.citizenship_id.toLowerCase().includes(searchLower) ||
                voterCard.location.toLowerCase().includes(searchLower)
            );
        });
    }

    async function createVoterCard() {
        try {
            const res = await fetch("http://localhost:3000/voter_card", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: JSON.stringify(newVoterCard),
            });
            if (!res.ok) throw new Error("Failed to create voter card");
            await fetchVoterCards();
            newVoterCard = {
                voter_card_id: "",
                citizenship_id: "",
                location: "",
            };
            newVoterCardModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    function openEditModal(voterCard) {
        editingVoterCard = { ...voterCard };
        editVoterCardModal.showModal();
    }

    async function updateVoterCard() {
        if (!editingVoterCard) return;
        try {
            const res = await fetch(
                `http://localhost:3000/voter_card/${editingVoterCard.voter_card_id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                    body: JSON.stringify(editingVoterCard),
                },
            );
            if (!res.ok) throw new Error("Failed to update voter card");
            await fetchVoterCards();
            editingVoterCard = null;
            editVoterCardModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    async function deleteVoterCard(id) {
        if (!confirm("Delete this voter card?")) return;
        try {
            const res = await fetch(`http://localhost:3000/voter_card/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete voter card");
            await fetchVoterCards();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchVoterCards);

    // Watch for search changes
    $: search, filterVoterCards();
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Manage Voter Cards
    </h1>
    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                placeholder="Search Voter Cards…"
                bind:value={search}
            />
        </div>
        <button
            class="btn btn-primary ml-4 rounded-lg"
            on:click={() => newVoterCardModal.showModal()}
        >
            + New Voter Card
        </button>
    </section>

    <dialog
        bind:this={newVoterCardModal}
        id="new_voter_card_modal"
        class="modal"
    >
        <div class="modal-box">
            <h3 class="font-bold text-lg">New Voter Card</h3>
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered"
                    placeholder="Voter Card ID"
                    bind:value={newVoterCard.voter_card_id}
                />
                <input
                    class="input input-bordered"
                    placeholder="Citizenship ID"
                    bind:value={newVoterCard.citizenship_id}
                />
                <input
                    class="input input-bordered"
                    placeholder="Location"
                    bind:value={newVoterCard.location}
                />
            </div>
            <div class="modal-action">
                <form method="dialog" class="flex gap-2">
                    <button
                        type="button"
                        class="btn btn-primary rounded-lg"
                        on:click={createVoterCard}>Create</button
                    >
                    <button class="btn btn-outline rounded-lg">Cancel</button>
                </form>
            </div>
        </div>
    </dialog>

    <dialog
        bind:this={editVoterCardModal}
        id="edit_voter_card_modal"
        class="modal"
    >
        <div class="modal-box">
            <h3 class="font-bold text-lg">Edit Voter Card</h3>
            {#if editingVoterCard}
                <div class="py-2 flex flex-col gap-2">
                    <label class="label">
                        <span class="label-text">Voter Card ID</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingVoterCard.voter_card_id}
                    />
                    <label class="label">
                        <span class="label-text">Citizenship ID</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingVoterCard.citizenship_id}
                    />
                    <label class="label">
                        <span class="label-text">Location</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingVoterCard.location}
                    />
                </div>
                <div class="modal-action">
                    <form method="dialog" class="flex gap-2">
                        <button
                            type="button"
                            class="btn btn-warning"
                            on:click={updateVoterCard}>Update</button
                        >
                        <button class="btn">Cancel</button>
                    </form>
                </div>
            {/if}
        </div>
    </dialog>

    {#if loading}
        <p>Loading voter cards...</p>
    {:else if error}
        <p class="text-red-500">{error}</p>
    {:else}
        <div class="overflow-x-auto">
            <table class="table w-full">
                <thead>
                    <tr>
                        <th>Voter Card ID</th>
                        <th>Citizenship ID</th>
                        <th>Location</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {#each filteredVoterCards as voterCard}
                        <tr>
                            <td>{voterCard.voter_card_id}</td>
                            <td>{voterCard.citizenship_id}</td>
                            <td>{voterCard.location}</td>
                            <td class="flex gap-2">
                                <button
                                    class="btn btn-sm btn-warning"
                                    on:click={() => openEditModal(voterCard)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-sm btn-error"
                                    on:click={() =>
                                        deleteVoterCard(
                                            voterCard.voter_card_id,
                                        )}>Delete</button
                                >
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    {/if}
</section>

<style>
    .modal-box {
        max-width: 560px;
        width: 90%;
    }
</style>
