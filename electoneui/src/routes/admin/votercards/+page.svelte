<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let voterCards = [];
    let filteredVoterCards = [];
    let loading = true;
    let error = "";
    let search = "";
    let editMessage = "";

    let offset = 0;
    let limit = 10;
    let hasMore = true;

    let newVoterCard = {
        voter_card_id: "",
        citizenship_id: "",
        location: "",
        voting_privileges: "",
    };

    let editingVoterCard = null;
    let newVoterCardModal;
    let editVoterCardModal;

    function isValidVoterCardId(voterCardId) {
        const regex = /^\d{10}$/;
        return regex.test(voterCardId);
    }

    function isValidCitizenshipId(citizenshipId) {
        return /^\d{2}-\d{2}-\d{2}-\d{5}$/.test(citizenshipId);
    }

    async function fetchVoterCards() {
        loading = true;
        try {
            const res = await fetch(
                `http://localhost:3000/voter_card?limit=${limit}&offset=${offset}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) {
                const errText = await res.json();
                throw new Error(errText.error);
            }
            voterCards = await res.json();
            filteredVoterCards = voterCards; // Initially, no filtering
            hasMore = voterCards.length === limit;
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    function nextPage() {
        offset += limit;
        fetchVoterCards();
    }

    function prevPage() {
        if (offset >= limit) {
            offset -= limit;
        } else {
            offset = 0;
        }
        fetchVoterCards();
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
        editMessage = "";

        // Check for empty fields
        if (
            !newVoterCard.voter_card_id ||
            !newVoterCard.citizenship_id ||
            !newVoterCard.location ||
            !newVoterCard.voting_privileges
        ) {
            editMessage = "All fields are required.";
            return;
        }

        // Validate fields
        if (!isValidVoterCardId(newVoterCard.voter_card_id)) {
            editMessage = "Invalid Voter Card ID. Must be 10 digits.";
            return;
        }

        if (!isValidCitizenshipId(newVoterCard.citizenship_id)) {
            editMessage = "Invalid Citizenship ID. Format: XX-XX-XX-XXXXX.";
            return;
        }

        try {
            const res = await fetch("http://localhost:3000/voter_card", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: JSON.stringify(newVoterCard),
            });
            if (!res.ok) {
                const errText = await res.json();
                throw new Error(errText.error);
            }

            await fetchVoterCards();
            newVoterCard = {
                voter_card_id: "",
                citizenship_id: "",
                location: "",
                voting_privileges: "",
            };
            newVoterCardModal.close();
        } catch (err) {
            editMessage = err.message;
        }
    }

    function openEditModal(voterCard) {
        editingVoterCard = { ...voterCard };
        editVoterCardModal.showModal();
    }

    async function updateVoterCard() {
        editMessage = ""; // Reset message before submitting

        // Check for empty fields
        if (
            !editingVoterCard.voter_card_id ||
            !editingVoterCard.citizenship_id ||
            !editingVoterCard.location ||
            !editingVoterCard.voting_privileges
        ) {
            editMessage = "All fields are required.";
            return;
        }

        // Validate fields
        if (!isValidVoterCardId(editingVoterCard.voter_card_id)) {
            editMessage = "Invalid Voter Card ID. Must be 10 digits.";
            return;
        }

        if (!isValidCitizenshipId(editingVoterCard.citizenship_id)) {
            editMessage = "Invalid Citizenship ID. Format: XX-XX-XX-XXXXX.";
            return;
        }

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
            console.log(res);

            if (!res.ok) {
                const errText = await res.json();
                throw new Error(errText.error);
            }
            await fetchVoterCards();
            editingVoterCard = null;
            editVoterCardModal.close();
        } catch (err) {
            editMessage = err.message;
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
            {#if editMessage}
                <p class="text-red-500 text-sm">{editMessage}</p>
            {/if}
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered w-full"
                    placeholder="Voter Card ID (XXXXXXXXXX)"
                    bind:value={newVoterCard.voter_card_id}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Citizenship ID (XX-XX-XX-XXXXX)"
                    bind:value={newVoterCard.citizenship_id}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Location"
                    bind:value={newVoterCard.location}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Additional Privileges"
                    bind:value={newVoterCard.voting_privileges}
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
                {#if editMessage}
                    <p class="text-red-500 text-sm">{editMessage}</p>
                {/if}
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
                    <label class="label">
                        <span class="label-text">Privileges</span>
                    </label>
                    <input
                        class="input input-bordered w-full"
                        placeholder="Additional Privileges"
                        bind:value={editingVoterCard.voting_privileges}
                    />
                </div>
                <div class="modal-action">
                    <form method="dialog" class="flex gap-2">
                        <button
                            type="button"
                            class="btn btn-primary"
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
                        <th>Privileges</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {#each filteredVoterCards as voterCard}
                        <tr>
                            <td>{voterCard.voter_card_id}</td>
                            <td>{voterCard.citizenship_id}</td>
                            <td>{voterCard.location}</td>
                            <td>{voterCard.voting_privileges}</td>
                            <td class="flex gap-2">
                                <button
                                    class="btn btn-sm btn-ghost"
                                    on:click={() => openEditModal(voterCard)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-sm btn-error text-white"
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
        <div class="flex justify-between items-center mt-4">
            <button
                class="btn btn-sm"
                disabled={offset === 0}
                on:click={prevPage}
            >
                Previous
            </button>
            <button class="btn btn-sm" disabled={!hasMore} on:click={nextPage}>
                Next
            </button>
        </div>
    {/if}
</section>

<style>
    .modal-box {
        max-width: 560px;
        width: 90%;
    }
</style>
