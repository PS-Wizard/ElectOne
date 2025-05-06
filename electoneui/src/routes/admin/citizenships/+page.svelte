<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let citizens = [];
    let filteredCitizens = [];
    let loading = true;
    let error = "";
    let search = "";

    let newCitizen = {
        citizenship_id: "",
        full_name: "",
        date_of_birth: "",
        birth_place: "",
        permanent_address: "",
    };

    let editingCitizen = null;
    let newCitizenModal;
    let editCitizenModal;

    async function fetchCitizenships() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/citizen?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch citizenships");
            citizens = await res.json();
            filteredCitizens = citizens; // Initially, no filtering
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    function filterCitizenships() {
        filteredCitizens = citizens.filter((citizen) => {
            const searchLower = search.toLowerCase();
            return (
                citizen.citizenship_id.toLowerCase().includes(searchLower) ||
                citizen.full_name.toLowerCase().includes(searchLower) ||
                citizen.date_of_birth.toLowerCase().includes(searchLower) ||
                citizen.birth_place.toLowerCase().includes(searchLower) ||
                citizen.permanent_address.toLowerCase().includes(searchLower)
            );
        });
    }

    async function createCitizenship() {
        try {
            const res = await fetch("http://localhost:3000/citizen", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: JSON.stringify(newCitizen),
            });
            if (!res.ok) throw new Error("Failed to create citizenship");
            await fetchCitizenships();
            newCitizen = {
                citizenship_id: "",
                full_name: "",
                date_of_birth: "",
                birth_place: "",
                permanent_address: "",
            };
            newCitizenModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    function openEditModal(citizen) {
        editingCitizen = { ...citizen };
        editCitizenModal.showModal();
    }

    async function updateCitizenship() {
        if (!editingCitizen) return;
        try {
            const res = await fetch(
                `http://localhost:3000/citizen/${editingCitizen.citizenship_id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                    body: JSON.stringify(editingCitizen),
                },
            );
            if (!res.ok) throw new Error("Failed to update citizenship");
            await fetchCitizenships();
            editingCitizen = null;
            editCitizenModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    async function deleteCitizenship(id) {
        if (!confirm("Delete this citizenship?")) return;
        try {
            const res = await fetch(`http://localhost:3000/citizen/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete citizenship");
            await fetchCitizenships();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchCitizenships);

    // Watch for search changes
    $: search, filterCitizenships();
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Manage Citizenships
    </h1>
    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                placeholder="Search Citizenships…"
                bind:value={search}
            />
        </div>
        <button
            class="btn btn-primary ml-4 rounded-lg"
            on:click={() => newCitizenModal.showModal()}
        >
            + New Citizenship
        </button>
    </section>

    <dialog
        bind:this={newCitizenModal}
        id="new_citizenship_modal"
        class="modal"
    >
        <div class="modal-box">
            <h3 class="font-bold text-lg">New Citizenship</h3>
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered"
                    placeholder="Citizenship ID"
                    bind:value={newCitizen.citizenship_id}
                />
                <input
                    class="input input-bordered"
                    placeholder="Full Name"
                    bind:value={newCitizen.full_name}
                />
                <input
                    class="input input-bordered"
                    placeholder="Date of Birth"
                    bind:value={newCitizen.date_of_birth}
                />
                <input
                    class="input input-bordered"
                    placeholder="Birth Place"
                    bind:value={newCitizen.birth_place}
                />
                <input
                    class="input input-bordered"
                    placeholder="Permanent Address"
                    bind:value={newCitizen.permanent_address}
                />
            </div>
            <div class="modal-action">
                <form method="dialog" class="flex gap-2">
                    <button
                        type="button"
                        class="btn btn-primary rounded-lg"
                        on:click={createCitizenship}>Create</button
                    >
                    <button class="btn btn-outline rounded-lg">Cancel</button>
                </form>
            </div>
        </div>
    </dialog>

    <dialog
        bind:this={editCitizenModal}
        id="edit_citizenship_modal"
        class="modal"
    >
        <div class="modal-box">
            <h3 class="font-bold text-lg">Edit Citizenship</h3>
            {#if editingCitizen}
                <div class="py-2 flex flex-col gap-2">
                    <label class="label">
                        <span class="label-text">Citizenship ID</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingCitizen.citizenship_id}
                    />
                    <label class="label">
                        <span class="label-text">Full Name</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingCitizen.full_name}
                    />
                    <label class="label">
                        <span class="label-text">Date of Birth</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingCitizen.date_of_birth}
                    />
                    <label class="label">
                        <span class="label-text">Birth Place</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingCitizen.birth_place}
                    />
                    <label class="label">
                        <span class="label-text">Permanent Address</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingCitizen.permanent_address}
                    />
                </div>
                <div class="modal-action">
                    <form method="dialog" class="flex gap-2">
                        <button
                            type="button"
                            class="btn btn-warning"
                            on:click={updateCitizenship}>Update</button
                        >
                        <button class="btn">Cancel</button>
                    </form>
                </div>
            {/if}
        </div>
    </dialog>

    {#if loading}
        <p>Loading citizenships...</p>
    {:else if error}
        <p class="text-red-500">{error}</p>
    {:else}
        <div class="overflow-x-auto">
            <table class="table w-full">
                <thead>
                    <tr>
                        <th>Citizenship ID</th>
                        <th>Full Name</th>
                        <th>Date of Birth</th>
                        <th>Birth Place</th>
                        <th>Permanent Address</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {#each filteredCitizens as citizen}
                        <tr>
                            <td>{citizen.citizenship_id}</td>
                            <td>{citizen.full_name}</td>
                            <td>{citizen.date_of_birth}</td>
                            <td>{citizen.birth_place}</td>
                            <td>{citizen.permanent_address}</td>
                            <td class="flex gap-2">
                                <button
                                    class="btn btn-sm btn-warning"
                                    on:click={() => openEditModal(citizen)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-sm btn-error"
                                    on:click={() =>
                                        deleteCitizenship(
                                            citizen.citizenship_id,
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
