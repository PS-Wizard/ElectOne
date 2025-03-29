<script>
    import { onMount } from "svelte";
    import { GetUsers, DeleteUser, UpdateUser, CreateUser } from "./api.js";

    let users = [];
    let newUser = { citizenID: "", password: "", tags: "", phoneNumber: "" };
    let editingUser = null;

    async function loadUsers() {
        users = await GetUsers();
    }

    async function removeUser(id) {
        await DeleteUser(id);
        loadUsers();
    }

    async function addUser() {
        await CreateUser(newUser);
        newUser = { citizenID: "", password: "", tags: "", phoneNumber: "" };
        loadUsers();
    }

    async function saveUser() {
        if (editingUser) {
            await UpdateUser(editingUser.userID, editingUser);
            editingUser = null;
            loadUsers();
        }
    }

    onMount(loadUsers);
</script>

<main class="flex w-full justify-center items-center flex-col">
    <h1>Users</h1>
    <ul>
        {#each users as user}
            <li>
                {#if editingUser?.userID === user.userID}
                    <input
                        bind:value={editingUser.citizenID}
                        placeholder="Citizen ID"
                    />
                    <input
                        bind:value={editingUser.password}
                        placeholder="Password"
                        type="password"
                    />
                    <input bind:value={editingUser.tags} placeholder="Tags" />
                    <input
                        bind:value={editingUser.phoneNumber}
                        placeholder="Phone Number"
                    />
                    <button on:click={saveUser}>Save</button>
                {:else}
                    {user.citizenID} - {user.tags} - {user.phoneNumber}
                    <button on:click={() => (editingUser = { ...user })}
                        >Edit</button
                    >
                    <button on:click={() => removeUser(user.userID)}
                        >Delete</button
                    >
                {/if}
            </li>
        {/each}
    </ul>
    <h2>Add User</h2>
    <input bind:value={newUser.citizenID} placeholder="Citizen ID" />
    <input
        bind:value={newUser.password}
        placeholder="Password"
        type="password"
    />
    <input bind:value={newUser.tags} placeholder="Tags" />
    <input bind:value={newUser.phoneNumber} placeholder="Phone Number" />
    <button on:click={addUser}>Add</button>
</main>
