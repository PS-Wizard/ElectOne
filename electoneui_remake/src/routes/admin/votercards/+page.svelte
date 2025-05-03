<script>
    import { onMount } from "svelte";

    let voterCards = [];
    let loading = true;
    let error = "";

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
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch voter cards");
            voterCards = await res.json();
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    async function createVoterCard() {
        try {
            const res = await fetch("http://localhost:3000/voter_card", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
                body: JSON.stringify(newVoterCard),
            });
            if (!res.ok) throw new Error("Failed to create voter card");
            await fetchVoterCards();
            newVoterCard = {
                voter_card_id: "",
                citizenship_id: "",
                location: "",
            }; // Reset form
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
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
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
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete voter card");
            await fetchVoterCards();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchVoterCards);
</script>

<button
    class="btn btn-primary mb-4"
    on:click={() => newVoterCardModal.showModal()}>New Voter Card</button
>

<dialog bind:this={newVoterCardModal} id="new_voter_card_modal" class="modal">
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
                    class="btn btn-success"
                    on:click={createVoterCard}>Create</button
                >
                <button class="btn">Cancel</button>
            </form>
        </div>
    </div>
</dialog>

<dialog bind:this={editVoterCardModal} id="edit_voter_card_modal" class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">Edit Voter Card</h3>
        {#if editingVoterCard}
            <div class="py-2 flex flex-col gap-2">
                <label class="label">
                    <span class="label-text">Voter Card ID</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
                    value={editingVoterCard.voter_card_id}
                    readonly
                />
                <label class="label">
                    <span class="label-text">Citizenship ID</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
                    bind:value={editingVoterCard.citizenship_id}
                />
                <label class="label">
                    <span class="label-text">Location</span>
                </label>
                <input
                    type="text"
                    class="input input-bordered"
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
                {#each voterCards as card}
                    <tr>
                        <td>{card.voter_card_id}</td>
                        <td>{card.citizenship_id}</td>
                        <td>{card.location}</td>
                        <td class="flex gap-2">
                            <button
                                class="btn btn-sm btn-warning"
                                on:click={() => openEditModal(card)}
                                >Edit</button
                            >
                            <button
                                class="btn btn-sm btn-error"
                                on:click={() =>
                                    deleteVoterCard(card.voter_card_id)}
                                >Delete</button
                            >
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
{/if}
