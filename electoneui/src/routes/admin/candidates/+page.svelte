<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let offset = 0;
    let limit = 10;
    let hasMore = true;

    let candidates = [];
    let loading = true;
    let error = "";
    let search = "";
    let selectedPhoto = "";
    let createMessage = "";
    let newCandidate = {
        citizen_id: "",
        election_id: "",
        candidate_bio: "",
        candidate_post: "",
        candidate_photo: null,
        candidate_party: "",
        candidate_name: "",
    };

    let editingCandidate = null;
    let editedPhoto = null;

    let newCandidateModal;
    let editCandidateModal;

    function isValidCitizenship(id) {
        return /^\d{2}-\d{2}-\d{2}-\d{5}$/.test(id);
    }

    function isValidFormImageOptional(candidate) {
        if (
            !candidate.candidate_name ||
            !candidate.candidate_party ||
            !candidate.citizen_id ||
            !candidate.election_id ||
            !candidate.candidate_bio ||
            !candidate.candidate_post
        ) {
            return { valid: false, message: "All fields must be filled in" };
        }
        if (!isValidCitizenship(candidate.citizen_id)) {
            return { valid: false, message: "Invalid Citizenship ID format" };
        }
        return { valid: true };
    }

    // Validate all fields
    function isValidForm(candidate) {
        if (
            !candidate.candidate_name ||
            !candidate.candidate_party ||
            !candidate.citizen_id ||
            !candidate.election_id ||
            !candidate.candidate_bio ||
            !candidate.candidate_post ||
            !candidate.candidate_photo
        ) {
            return { valid: false, message: "All fields must be filled in" };
        }
        if (!isValidCitizenship(candidate.citizen_id)) {
            return { valid: false, message: "Invalid Citizenship ID format" };
        }
        return { valid: true };
    }

    function openEditModal(candidate) {
        editingCandidate = { ...candidate };
        editedPhoto = null;
        editCandidateModal.showModal();
    }
    async function fetchCandidates() {
        loading = true;
        try {
            const res = await fetch(
                `http://localhost:3000/candidate?limit=${limit}&offset=${offset}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch candidates");
            candidates = await res.json();
            hasMore = candidates.length === limit;
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    function nextPage() {
        offset += limit;
        fetchCandidates();
    }

    function prevPage() {
        if (offset >= limit) {
            offset -= limit;
        } else {
            offset = 0;
        }
        fetchCandidates();
    }

    async function createCandidate() {
        const validation = isValidForm(newCandidate);
        if (!validation.valid) {
            createMessage = validation.message;
            return;
        }

        const formData = new FormData();
        formData.append("citizenship_id", newCandidate.citizen_id);
        formData.append("election_id", newCandidate.election_id);
        formData.append("candidate_name", newCandidate.candidate_name);
        formData.append("candidate_party", newCandidate.candidate_party);
        formData.append("candidate_bio", newCandidate.candidate_bio);
        formData.append("candidate_post", newCandidate.candidate_post);
        formData.append("candidate_photo", newCandidate.candidate_photo);

        try {
            const res = await fetch("http://localhost:3000/candidate", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
                body: formData,
            });
            if (!res.ok) {
                const errText = await res.json();
                throw new Error(errText.error);
            }
            await fetchCandidates();
            newCandidate = {
                citizen_id: "",
                election_id: "",
                candidate_bio: "",
                candidate_post: "",
                candidate_photo: null,
            };
            newCandidateModal.close();
        } catch (err) {
            createMessage = err;
        }
    }

    async function updateCandidate() {
        if (!editingCandidate) return;

        const validation = isValidFormImageOptional(editingCandidate);
        if (!validation.valid) {
            createMessage = validation.message;
            return;
        }

        const formData = new FormData();
        formData.append("citizenship_id", editingCandidate.citizen_id);
        formData.append("election_id", editingCandidate.election_id);
        formData.append("candidate_bio", editingCandidate.candidate_bio);
        formData.append("candidate_post", editingCandidate.candidate_post);
        formData.append("candidate_name", editingCandidate.candidate_name);
        formData.append("candidate_party", editingCandidate.candidate_party);
        if (editedPhoto) {
            formData.append("candidate_photo", editedPhoto);
        }

        try {
            const res = await fetch(
                `http://localhost:3000/candidate/${editingCandidate.candidate_id}`,
                {
                    method: "PUT",
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                    body: formData,
                },
            );
            if (!res.ok) {
                const errText = await res.json();
                throw new Error(errText.error);
            }
            await fetchCandidates();
            editingCandidate = null;
            editedPhoto = null;
            editCandidateModal.close();
        } catch (err) {
            createMessage = err;
        }
    }

    async function deleteCandidate(id) {
        if (!confirm("Delete this candidate?")) return;
        try {
            const res = await fetch(`http://localhost:3000/candidate/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete candidate");
            await fetchCandidates();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchCandidates);
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Manage Candidates
    </h1>

    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                placeholder="Search candidates…"
                bind:value={search}
            />
        </div>
        <button
            class="btn btn-primary ml-4 rounded-lg"
            on:click={() => newCandidateModal.showModal()}
        >
            + New Candidate
        </button>
    </section>

    <section class="flex justify-center m-4">
        {#if loading}
            <p>Loading candidates...</p>
        {:else if error}
            <p class="text-red-500">{error}</p>
        {:else}
            <div
                class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100 w-full"
            >
                <table class="table w-full">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <!-- New column -->
                            <th>Citizen ID</th>
                            <th>Election ID</th>
                            <th>Profile</th>
                            <th>Party</th>
                            <!-- Optional, might be helpful -->
                            <th>Bio</th>
                            <th>Post</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each candidates as c}
                            {#if c.citizen_id
                                .toLowerCase()
                                .includes(search.toLowerCase()) || c.election_id
                                    .toString()
                                    .includes(search)}
                                <tr>
                                    <td>{c.candidate_id}</td>
                                    <td>{c.candidate_name}</td>
                                    <!-- New data cell -->
                                    <td>{c.citizen_id}</td>
                                    <td>{c.election_id}</td>
                                    <td>
                                        <img
                                            src={c.candidate_photo}
                                            alt="Candidate Photo"
                                            class="w-12 h-12 object-cover rounded-full cursor-pointer hover:scale-105 transition"
                                            on:click={() =>
                                                (selectedPhoto = c.candidate_photo)}
                                        />
                                    </td>
                                    <td>{c.candidate_party}</td>
                                    <!-- Optional -->
                                    <td>{c.candidate_bio}</td>
                                    <td>{c.candidate_post}</td>
                                    <td class="flex gap-2">
                                        <button
                                            class="btn btn-sm btn-ghost"
                                            on:click={() => openEditModal(c)}
                                        >
                                            Edit
                                        </button>
                                        <button
                                            class="btn btn-sm btn-error text-white"
                                            on:click={() =>
                                                deleteCandidate(c.candidate_id)}
                                        >
                                            Delete
                                        </button>
                                    </td>
                                </tr>
                            {/if}
                        {/each}
                    </tbody>
                </table>
            </div>
        {/if}
    </section>
    <div class="flex justify-between items-center mt-4">
        <button class="btn btn-sm" disabled={offset === 0} on:click={prevPage}>
            Previous
        </button>
        <button class="btn btn-sm" disabled={!hasMore} on:click={nextPage}>
            Next
        </button>
    </div>
</section>

<dialog bind:this={newCandidateModal} class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">Create New Candidate</h3>
        <div class="py-2 flex flex-col gap-2">
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Full Name"
                bind:value={newCandidate.candidate_name}
                required
            />
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Party"
                bind:value={newCandidate.candidate_party}
                required
            />
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Citizen ID"
                bind:value={newCandidate.citizen_id}
                required
            />
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Election ID"
                bind:value={newCandidate.election_id}
                required
            />
            <label for="new_candidate_photo" class="label">
                <span class="label-text">Profile Photo</span>
            </label>
            <input
                type="file"
                id="new_candidate_photo"
                class="file-input file-input-bordered w-full rounded-lg"
                accept="image/*"
                on:change={(e) =>
                    (newCandidate.candidate_photo = e.target.files[0])}
                required
            />
            <textarea
                class="textarea textarea-bordered w-full rounded-lg"
                placeholder="Bio"
                bind:value={newCandidate.candidate_bio}
                required
            />
            <input
                class="input input-bordered w-full rounded-lg"
                placeholder="Post"
                bind:value={newCandidate.candidate_post}
                required
            />
        </div>
        {#if createMessage}
            <p class="text-sm text-red-700">{createMessage}</p>
        {/if}
        <div class="modal-action">
            <form method="dialog" class="flex gap-2">
                <button
                    type="button"
                    class="btn btn-primary rounded-lg"
                    on:click={createCandidate}>Create</button
                >
                <button class="btn btn-outline rounded-lg">Cancel</button>
            </form>
        </div>
    </div>
</dialog>

<dialog bind:this={editCandidateModal} class="modal">
    <div class="modal-box">
        <h3 class="font-bold text-lg">Edit Candidate</h3>
        {#if editingCandidate}
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered w-full border-black"
                    value={editingCandidate.candidate_id}
                    readonly
                />
                <input
                    class="input input-bordered w-full rounded-lg"
                    placeholder="Full Name"
                    bind:value={editingCandidate.candidate_name}
                />
                <input
                    class="input input-bordered w-full rounded-lg"
                    placeholder="Party"
                    bind:value={editingCandidate.candidate_party}
                />
                <input
                    class="input input-bordered w-full rounded-lg"
                    placeholder="Citizen ID"
                    bind:value={editingCandidate.citizen_id}
                />
                <input
                    class="input input-bordered w-full rounded-lg"
                    placeholder="Election ID"
                    bind:value={editingCandidate.election_id}
                />
                <label for="edit_candidate_photo" class="label">
                    <span class="label-text"
                        >Update Profile Photo (Optional)</span
                    >
                </label>
                <input
                    type="file"
                    id="edit_candidate_photo"
                    class="file-input file-input-bordered w-full rounded-lg"
                    accept="image/*"
                    on:change={(e) => (editedPhoto = e.target.files[0])}
                />
                <textarea
                    class="textarea textarea-bordered w-full rounded-lg"
                    placeholder="Bio"
                    bind:value={editingCandidate.candidate_bio}
                />
                <input
                    class="input input-bordered w-full rounded-lg"
                    placeholder="Post"
                    bind:value={editingCandidate.candidate_post}
                />
            </div>
            {#if createMessage}
                <p class="text-red-700 text-sm">{createMessage}</p>
            {/if}
            <div class="modal-action">
                <form method="dialog" class="flex gap-2">
                    <button
                        type="button"
                        class="btn btn-primary"
                        on:click={updateCandidate}>Update</button
                    >
                    <button class="btn">Cancel</button>
                </form>
            </div>
        {/if}
    </div>
</dialog>

{#if selectedPhoto}
    <div
        class="fixed inset-0 bg-black/80 flex items-center justify-center z-50"
    >
        <div class="relative">
            <img
                src={selectedPhoto}
                class="max-w-full max-h-screen rounded-lg shadow-lg"
            />
            <button
                class="absolute top-2 right-2 btn btn-sm btn-circle btn-error"
                on:click={() => (selectedPhoto = null)}
            >
                ✕
            </button>
        </div>
    </div>
{/if}

<style>
    .modal-box {
        max-width: 560px;
        width: 90%;
    }
</style>
