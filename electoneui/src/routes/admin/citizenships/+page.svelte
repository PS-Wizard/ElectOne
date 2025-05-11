<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let citizens = [];
    let filteredCitizens = [];
    let loading = true;
    let error = "";
    let search = "";

    let createMessage = "";
    let editMessage = "";

    let offset = 0;
    let limit = 10;
    let hasMore = true;

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

    function isValidCitizenship(id) {
        return /^\d{2}-\d{2}-\d{2}-\d{5}$/.test(id);
    }

    async function fetchCitizenships() {
        loading = true;
        try {
            const res = await fetch(
                `http://localhost:3000/citizen?limit=${limit}&offset=${offset}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch citizenships");
            citizens = await res.json();
            filteredCitizens = citizens;
            hasMore = citizens.length === limit;
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    function nextPage() {
        offset += limit;
        fetchCitizenships();
    }

    function prevPage() {
        if (offset >= limit) {
            offset -= limit;
        } else {
            offset = 0;
        }
        fetchCitizenships();
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
        createMessage = "";

        const {
            citizenship_id,
            full_name,
            date_of_birth,
            birth_place,
            permanent_address,
        } = newCitizen;

        if (
            !citizenship_id ||
            !full_name ||
            !date_of_birth ||
            !birth_place ||
            !permanent_address
        ) {
            createMessage = "Please fill in all fields.";
            return;
        }

        if (!isValidCitizenship(citizenship_id)) {
            createMessage =
                "Invalid Citizenship ID. Format must be XX-XX-XX-XXXXX";
            return;
        }

        try {
            const res = await fetch("http://localhost:3000/citizen", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: JSON.stringify(newCitizen),
            });

            if (!res.ok) {
                const err = await res.json();
                createMessage = err.message || "Failed to create citizenship.";
                return;
            }

            await fetchCitizenships();
            newCitizen = {
                citizenship_id: "",
                full_name: "",
                date_of_birth: "",
                birth_place: "",
                permanent_address: "",
            };
            createMessage = "";
            newCitizenModal.close();
        } catch (err) {
            createMessage = err.message;
        }
    }

    function openEditModal(citizen) {
        editingCitizen = { ...citizen };
        editCitizenModal.showModal();
    }

    async function updateCitizenship() {
        editMessage = "";
        if (!editingCitizen) return;

        const {
            citizenship_id,
            full_name,
            date_of_birth,
            birth_place,
            permanent_address,
        } = editingCitizen;

        if (
            !citizenship_id ||
            !full_name ||
            !date_of_birth ||
            !birth_place ||
            !permanent_address
        ) {
            editMessage = "Please fill in all fields.";
            return;
        }

        if (!isValidCitizenship(citizenship_id)) {
            editMessage =
                "Invalid Citizenship ID. Format must be XX-XX-XX-XXXXX";
            return;
        }

        try {
            const res = await fetch(
                `http://localhost:3000/citizen/${citizenship_id}`,
                {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                    body: JSON.stringify(editingCitizen),
                },
            );

            if (!res.ok) {
                const err = await res.json();
                editMessage = err.message || "Failed to update citizenship.";
                return;
            }

            await fetchCitizenships();
            editingCitizen = null;
            editCitizenModal.close();
        } catch (err) {
            editMessage = err.message;
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
                    type="text"
                    class="input input-bordered w-full"
                    placeholder="XX-XX-XX-XXXXX"
                    required
                    bind:value={newCitizen.citizenship_id}
                />
                <input
                    class="input input-bordered w-full"
                    required
                    type="text"
                    placeholder="Jhon Doe"
                    bind:value={newCitizen.full_name}
                />
                <input
                    type="date"
                    class="input input-bordered w-full"
                    required
                    placeholder="Date of Birth"
                    bind:value={newCitizen.date_of_birth}
                />
                <input
                    class="input input-bordered w-full"
                    required
                    placeholder="Kathmandu, Nepal"
                    bind:value={newCitizen.birth_place}
                />
                <input
                    class="input input-bordered w-full"
                    required
                    placeholder="Kathmandu, Nepal"
                    bind:value={newCitizen.permanent_address}
                />
            </div>
            {#if createMessage}
                <p class="text-red-500 mt-2 text-sm">{createMessage}</p>
            {/if}
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
                        type="date"
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
                {#if editMessage}
                    <p class="text-red-500 mt-2 text-sm">{editMessage}</p>
                {/if}

                <div class="modal-action">
                    <form method="dialog" class="flex gap-2">
                        <button
                            type="button"
                            class="btn btn-primary"
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
                                    class="btn btn-sm btn-ghost"
                                    on:click={() => openEditModal(citizen)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-sm text-white btn-error"
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
