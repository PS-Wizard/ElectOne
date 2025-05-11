<script>
    import { onMount } from "svelte";
    import AdminNavbar from "../../../components/AdminNavbar.svelte";

    let users = [];
    let filteredUsers = [];
    let loading = true;
    let error = "";
    let creationMessage = "";
    let search = "";
    let selectedPhoto = "";

    function isValidCitizenship(id) {
        return /^\d{2}-\d{2}-\d{2}-\d{5}$/.test(id);
    }

    function isValidVoterCard(id) {
        return /^\d{10}$/.test(id);
    }

    function isValidPassword(pw) {
        return (
            pw.length >= 8 &&
            /[A-Z]/.test(pw) &&
            /\d/.test(pw) &&
            /[^A-Za-z0-9]/.test(pw)
        );
    }

    let newUser = {
        citizenship_id: "",
        voter_card_id: "",
        password: "",
        citizenship_front: null,
        citizenship_back: null,
        voter_card: null,
        selfie: null,
    };

    let editingUser = null;
    let newUserModal;
    let editUserModal;
    let editedPhotos = {
        citizenship_front: null,
        citizenship_back: null,
        voter_card: null,
        selfie: null,
    };

    async function fetchUsers() {
        loading = true;
        try {
            const res = await fetch(
                "http://localhost:3000/user?limit=100&offset=0",
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                    },
                },
            );
            if (!res.ok) throw new Error("Failed to fetch users");
            users = await res.json();
            filteredUsers = users; // Initially, no filtering
        } catch (err) {
            error = err.message;
        }
        loading = false;
    }

    function filterUsers() {
        filteredUsers = users.filter((user) => {
            const searchLower = search.toLowerCase();
            return (
                user.user_id.toString().toLowerCase().includes(searchLower) ||
                user.citizenship_id.toLowerCase().includes(searchLower) ||
                user.voter_card_id.toLowerCase().includes(searchLower) ||
                user.totp_secret.toLowerCase().includes(searchLower) ||
                (user.photos && user.photos.toLowerCase().includes(searchLower))
            );
        });
    }

    async function createUser() {
        if (
            !newUser.citizenship_front ||
            !newUser.citizenship_back ||
            !newUser.voter_card ||
            !newUser.selfie
        ) {
            creationMessage = "Please upload all 4 required photos.";
            return;
        }

        if (!isValidCitizenship(newUser.citizenship_id)) {
            creationMessage = "Invalid Citizenship ID format.";
            return;
        }

        if (!isValidVoterCard(newUser.voter_card_id)) {
            creationMessage = "Invalid Voter Card ID.";
            return;
        }

        if (!isValidPassword(newUser.password)) {
            creationMessage =
                "Password must be at least 8 chars, include a number, uppercase, and symbol.";
            return;
        }

        const formData = new FormData();
        formData.append("citizenship_id", newUser.citizenship_id);
        formData.append("voter_card_id", newUser.voter_card_id);
        formData.append("password", newUser.password);
        formData.append("photos", newUser.citizenship_front);
        formData.append("photos", newUser.citizenship_back);
        formData.append("photos", newUser.voter_card);
        formData.append("photos", newUser.selfie);

        try {
            const res = await fetch("http://localhost:3000/user", {
                method: "POST",
                body: formData,
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            const data = await res.json();

            if (!res.ok) {
                alert(
                    data.message ||
                        "Something went wrong during user creation.",
                );
                return;
            }

            alert(`User created successfully! ID: ${data.user_id}`);
            newUserModal.close();
            // Optionally reset the form
            newUser = {
                citizenship_id: "",
                voter_card_id: "",
                password: "",
                citizenship_front: null,
                citizenship_back: null,
                voter_card: null,
                selfie: null,
            };
            fetchUsers();
        } catch (err) {
            alert("Error creating user: " + err.message);
        }
    }

    function openEditModal(user) {
        editingUser = { ...user };
        editedPhotos = {
            citizenship_front: null,
            citizenship_back: null,
            voter_card: null,
            selfie: null,
        };
        editUserModal.showModal();
    }

    async function handleUpdateUser() {
        if (!editingUser) return;

        if (!isValidCitizenship(editingUser.citizenship_id)) {
            creationMessage = "Invalid Citizenship ID format.";
            return;
        }

        if (!isValidVoterCard(editingUser.voter_card_id)) {
            creationMessage = "Invalid Voter Card ID.";
            return;
        }

        if (!isValidPassword(editingUser.password)) {
            creationMessage =
                "Password must be at least 8 chars, include a number, uppercase, and symbol.";
            return;
        }

        const formData = new FormData();
        formData.append("citizenship_id", editingUser.citizenship_id);
        formData.append("voter_card_id", editingUser.voter_card_id);
        formData.append("password", editingUser.password);
        formData.append("totp_secret", editingUser.totp_secret);

        let hasNewPhotos = false;
        if (editedPhotos.citizenship_front) {
            formData.append("photos", editedPhotos.citizenship_front);
            hasNewPhotos = true;
        }
        if (editedPhotos.citizenship_back) {
            formData.append("photos", editedPhotos.citizenship_back);
            hasNewPhotos = true;
        }
        if (editedPhotos.voter_card) {
            formData.append("photos", editedPhotos.voter_card);
            hasNewPhotos = true;
        }
        if (editedPhotos.selfie) {
            formData.append("photos", editedPhotos.selfie);
            hasNewPhotos = true;
        }

        const headers = {
            Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
        };
        try {
            const res = await fetch(
                `http://localhost:3000/user/${editingUser.user_id}`,
                {
                    method: "PUT",
                    headers: headers,
                    body: formData,
                },
            );
            if (!res.ok) {
                const errorData = await res.json();
                throw new Error(errorData.message || "Failed to update user");
            }
            await fetchUsers();
            editingUser = null;
            editUserModal.close();
        } catch (err) {
            alert(err.message);
        }
    }

    async function deleteUser(id) {
        if (!confirm("Delete this user?")) return;
        try {
            const res = await fetch(`http://localhost:3000/user/${id}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("admin_token")}`,
                },
            });
            if (!res.ok) throw new Error("Failed to delete user");
            await fetchUsers();
        } catch (err) {
            alert(err.message);
        }
    }

    onMount(fetchUsers);

    // Watch for search changes
    $: search, filterUsers();
</script>

<AdminNavbar />
<section class="flex p-12 flex-col justify-center">
    <h1 class="text-center text-6xl uppercase font-medium tracking-wide">
        Manage Users
    </h1>
    <section class="flex justify-between items-center m-4">
        <div class="form-control w-full max-w-md">
            <input
                type="text"
                class="input rounded-lg w-full"
                placeholder="Search Users…"
                bind:value={search}
            />
        </div>
        <button
            class="btn btn-primary ml-4 rounded-lg"
            on:click={() => newUserModal.showModal()}
        >
            + New User
        </button>
    </section>

    <dialog bind:this={newUserModal} id="new_user_modal" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">New User</h3>
            <div class="py-2 flex flex-col gap-2">
                <input
                    class="input input-bordered w-full"
                    placeholder="Citizenship ID"
                    bind:value={newUser.citizenship_id}
                />
                <input
                    class="input input-bordered w-full"
                    placeholder="Voter Card ID"
                    bind:value={newUser.voter_card_id}
                />
                <input
                    type="password"
                    class="input input-bordered w-full"
                    placeholder="Password"
                    bind:value={newUser.password}
                />
                <label for="citizenship_front" class="label">
                    <span class="label-text">Citizenship ID (Front) </span>
                </label>
                <input
                    type="file"
                    id="citizenship_front"
                    class="file-input file-input-bordered w-full"
                    accept="image/*"
                    on:change={(e) =>
                        (newUser.citizenship_front = e.target.files
                            ? e.target.files.item(0)
                            : null)}
                />
                <label for="citizenship_back" class="label">
                    <span class="label-text">Citizenship ID (Back) </span>
                </label>
                <input
                    type="file"
                    id="citizenship_back"
                    class="file-input file-input-bordered w-full"
                    accept="image/*"
                    on:change={(e) =>
                        (newUser.citizenship_back = e.target.files
                            ? e.target.files.item(0)
                            : null)}
                />
                <label for="voter_card" class="label">
                    <span class="label-text">Voter Card </span>
                </label>
                <input
                    type="file"
                    id="voter_card"
                    class="file-input file-input-bordered w-full"
                    accept="image/*"
                    on:change={(e) =>
                        (newUser.voter_card = e.target.files
                            ? e.target.files.item(0)
                            : null)}
                />
                <label for="selfie" class="label">
                    <span class="label-text">Selfie </span>
                </label>
                <input
                    type="file"
                    id="selfie"
                    class="file-input file-input-bordered w-full"
                    accept="image/*"
                    on:change={(e) =>
                        (newUser.selfie = e.target.files
                            ? e.target.files.item(0)
                            : null)}
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
                        on:click={createUser}>Create</button
                    >
                    <button class="btn btn-outline rounded-lg">Cancel</button>
                </form>
            </div>
        </div>
    </dialog>

    <dialog bind:this={editUserModal} id="edit_user_modal" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">Edit User</h3>
            {#if editingUser}
                <div class="py-2 flex flex-col gap-2">
                    <label class="label">
                        <span class="label-text">Citizenship ID</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingUser.citizenship_id}
                    />
                    <label class="label">
                        <span class="label-text">Voter Card ID</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingUser.voter_card_id}
                    />
                    <label class="label">
                        <span class="label-text">Password</span>
                    </label>
                    <input
                        type="password"
                        class="input input-bordered w-full px-4 rounded-lg"
                        placeholder="Password"
                        bind:value={editingUser.password}
                    />
                    <label class="label">
                        <span class="label-text">TOTP Secret</span>
                    </label>
                    <input
                        type="text"
                        class="input input-bordered w-full px-4 rounded-lg"
                        bind:value={editingUser.totp_secret}
                    />

                    <label class="label">
                        <span class="label-text">Citizenship ID (Front)</span>
                    </label>
                    <input
                        type="file"
                        id="edit_citizenship_front"
                        class="file-input file-input-bordered w-full"
                        accept="image/*"
                        on:change={(e) =>
                            (editedPhotos.citizenship_front = e.target.files
                                ? e.target.files.item(0)
                                : null)}
                    />

                    <label class="label">
                        <span class="label-text">Citizenship ID (Back)</span>
                    </label>
                    <input
                        type="file"
                        id="edit_citizenship_back"
                        class="file-input file-input-bordered w-full"
                        accept="image/*"
                        on:change={(e) =>
                            (editedPhotos.citizenship_back = e.target.files
                                ? e.target.files.item(0)
                                : null)}
                    />

                    <label class="label">
                        <span class="label-text">Voter Card</span>
                    </label>
                    <input
                        type="file"
                        id="edit_voter_card"
                        class="file-input file-input-bordered w-full"
                        accept="image/*"
                        on:change={(e) =>
                            (editedPhotos.voter_card = e.target.files
                                ? e.target.files.item(0)
                                : null)}
                    />

                    <label class="label">
                        <span class="label-text">Selfie</span>
                    </label>
                    <input
                        type="file"
                        id="edit_selfie"
                        class="file-input file-input-bordered w-full"
                        accept="image/*"
                        on:change={(e) =>
                            (editedPhotos.selfie = e.target.files
                                ? e.target.files.item(0)
                                : null)}
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
                            on:click={handleUpdateUser}>Update</button
                        >
                        <button class="btn">Cancel</button>
                    </form>
                </div>
            {/if}
        </div>
    </dialog>

    {#if loading}
        <p>Loading users...</p>
    {:else if error}
        <p class="text-red-500">{error}</p>
    {:else}
        <div class="overflow-x-auto">
            <table class="table w-full">
                <thead>
                    <tr>
                        <th>User ID</th>
                        <th>Citizenship ID</th>
                        <th>Voter Card ID</th>
                        <th>TOTP Secret</th>
                        <th>Photos</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {#each filteredUsers as user}
                        <tr>
                            <td>{user.user_id}</td>
                            <td>{user.citizenship_id}</td>
                            <td>{user.voter_card_id}</td>
                            <td>{user.totp_secret}</td>
                            <td>
                                <div class="flex gap-1">
                                    {#each user.photos.split(",") as photo}
                                        <img
                                            src={`http://localhost:3000${photo}`}
                                            alt="Appeal Photo"
                                            class="w-12 h-12 object-cover rounded-full cursor-pointer hover:scale-105 transition"
                                            on:click={() =>
                                                (selectedPhoto = `http://localhost:3000${photo}`)}
                                        />
                                    {/each}
                                </div>
                            </td>

                            <td class="flex gap-2">
                                <button
                                    class="btn btn-sm btn-ghost "
                                    on:click={() => openEditModal(user)}
                                    >Edit</button
                                >
                                <button
                                    class="btn btn-sm btn-error text-white"
                                    on:click={() => deleteUser(user.user_id)}
                                    >Delete</button
                                >
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    {/if}
</section>

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
