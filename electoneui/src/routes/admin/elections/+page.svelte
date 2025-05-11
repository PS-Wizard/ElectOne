<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let elections = [];
    let filteredElections = [];
    let loading = true;
    let error = "";
    let search = "";

    let offset = 0;
    let limit = 10;
    let hasMore = true;

    let creationMessage = "";

    let newElection = {
        name: "",
        description: "",
        start_date: "",
        end_date: "",
        location: "",
        type: "",
    };

    let editingElection = null;
    let newElectionModal;
    let editElectionModal;

    async function fetchElections() {
        loading = true;
        try {
            const res = await fetch(
                `http://localhost:3000/election?limit=${limit}&offset=${offset}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch elections");
            elections = await res.json();
            filteredElections = elections; // Initially, no filtering
            hasMore = elections.length === limit;
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }
    function nextPage() {
        offset += limit;
        fetchElections();
    }

    function prevPage() {
        if (offset >= limit) {
            offset -= limit;
        } else {
            offset = 0;
        }
        fetchElections();
    }

    function filterElections() {
        filteredElections = elections.filter((election) => {
            const searchLower = search.toLowerCase();
            return (
                election.name.toLowerCase().includes(searchLower) ||
                election.description.toLowerCase().includes(searchLower) ||
                election.location.toLowerCase().includes(searchLower)
            );
        });
    }

    function validateElection(election) {
        // Basic validation to ensure no field is empty
        return (
            election.name.trim() !== "" &&
            election.description.trim() !== "" &&
            election.start_date !== "" &&
            election.end_date !== "" &&
            election.type !== "" &&
            election.location.trim() !== ""
        );
    }
    async function createElection() {
        if (!validateElection(newElection)) {
            creationMessage = "Please Fill In All The Fields!";
            return;
        }
        try {
            const res = await fetch("http://localhost:3000/election", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: JSON.stringify(newElection),
            });
            if (!res.ok) throw new Error("Failed to create election");
            await fetchElections();
            newElection = {
                name: "",
                description: "",
                start_date: "",
                end_date: "",
                location: "",
                type: "",
            };
            newElectionModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    function openEditModal(election) {
        editingElection = { ...election };
        editElectionModal.showModal();
    }

    async function updateElection() {
        if (!editingElection) return;
        if (!validateElection(editingElection)) {
            creationMessage = "Please Fill In All The Fields!";
            return;
        }
        try {
            const res = await fetch(
                `http://localhost:3000/election/${editingElection.election_id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                    body: JSON.stringify(editingElection),
                },
            );
            if (!res.ok) throw new Error("Failed to update election");
            await fetchElections();
            editingElection = null;
            editElectionModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    async function deleteElection(id) {
        if (!confirm("Delete this election?")) return;
        try {
            const res = await fetch(`http://localhost:3000/election/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete election");
            await fetchElections();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchElections);

    // Watch for search changes
    $: search, filterElections();
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Manage Elections
    </h1>
    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                placeholder="Search Elections…"
                bind:value={search}
            />
        </div>
        <button
            class="btn btn-primary ml-4 rounded-lg"
            on:click={() => newElectionModal.showModal()}
        >
            + New Election
        </button>
    </section>

    <dialog bind:this={newElectionModal} id="new_election_modal" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">New Election</h3>
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered w-full"
                    placeholder="Election Name"
                    required
                    bind:value={newElection.name}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Description"
                    required
                    bind:value={newElection.description}
                />
                <input
                    type="date"
                    class="input input-bordered w-full"
                    placeholder="Start Date"
                    required
                    bind:value={newElection.start_date}
                />
                <input
                    type="date"
                    class="input input-bordered w-full"
                    placeholder="End Date"
                    required
                    bind:value={newElection.end_date}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Location"
                    required
                    bind:value={newElection.location}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Type Of Election"
                    required
                    bind:value={newElection.type}
                />
            </div>
            {#if creationMessage}
                <p class="text-sm text-red-700">{creationMessage}</p>
            {/if}
            <div class="modal-action">
                <form method="dialog" class="flex gap-2">
                    <button
                        type="button"
                        class="btn btn-primary rounded-lg"
                        on:click={createElection}>Create</button
                    >
                    <button class="btn btn-outline rounded-lg">Cancel</button>
                </form>
            </div>
        </div>
    </dialog>

    <dialog
        bind:this={editElectionModal}
        id="edit_election_modal"
        class="modal"
    >
        <div class="modal-box">
            <h3 class="font-bold text-lg">Edit Election</h3>
            {#if editingElection}
                <div class="py-2 flex flex-col gap-2">
                    <label class="label">
                        <span class="label-text">Election Name</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingElection.name}
                    />
                    <label class="label">
                        <span class="label-text">Description</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingElection.description}
                    />
                    <label class="label">
                        <span class="label-text">Start Date</span>
                    </label>
                    <input
                        type="date"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingElection.start_date}
                    />
                    <label class="label">
                        <span class="label-text">End Date</span>
                    </label>
                    <input
                        type="date"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingElection.end_date}
                    />
                    <label class="label">
                        <span class="label-text">Location</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingElection.location}
                    />
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingElection.type}
                    />
                </div>
                {#if creationMessage}
                    <p class="text-sm text-red-700">{creationMessage}</p>
                {/if}
                <div class="modal-action">
                    <form method="dialog" class="flex gap-2">
                        <button
                            type="button"
                            class="btn btn-primary"
                            on:click={updateElection}>Update</button
                        >
                        <button class="btn">Cancel</button>
                    </form>
                </div>
            {/if}
        </div>
    </dialog>

    {#if loading}
        <p>Loading elections...</p>
    {:else if error}
        <p class="text-red-500">{error}</p>
    {:else}
        <div class="overflow-x-auto">
            <table class="table w-full">
                <thead>
                    <tr>
                        <th>Election ID</th>
                        <th>Name</th>
                        <th>Description</th>
                        <th>Start Date</th>
                        <th>End Date</th>
                        <th>Location</th>
                        <th>Type</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {#each filteredElections as election}
                        <tr>
                            <td>{election.election_id}</td>
                            <td>{election.name}</td>
                            <td>{election.description}</td>
                            <td>{election.start_date}</td>
                            <td>{election.end_date}</td>
                            <td>{election.location}</td>
                            <td>{election.type}</td>
                            <td class="flex gap-2">
                                <button
                                    class="btn btn-sm btn-ghost"
                                    on:click={() => openEditModal(election)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-sm btn-error text-white"
                                    on:click={() =>
                                        deleteElection(election.election_id)}
                                    >Delete</button
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
