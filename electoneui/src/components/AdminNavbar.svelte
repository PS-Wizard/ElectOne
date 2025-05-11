<script>
    import { goto } from "$app/navigation";

    const links = [
        { label: "Appeals", path: "/admin/appeals" },
        { label: "Candidates", path: "/admin/candidates" },
        { label: "Citizenships", path: "/admin/citizenships" },
        { label: "Elections", path: "/admin/elections" },
        { label: "Users", path: "/admin/users" },
        { label: "Voter Cards", path: "/admin/votercards" },
    ];

    function logout() {
        // handle your logout logic here
        localStorage.removeItem("admin_token");
        goto("/admin");
    }
</script>

<nav class="navbar bg-base-200 px-6 shadow">
    <div class="flex-1">
        <ul class="menu menu-horizontal px-1 gap-2">
            {#each links as link}
                <li>
                    <a href={link.path} class="font-medium">{link.label}</a>
                </li>
            {/each}
        </ul>
    </div>

    <div class="flex-none">
        <button on:click={logout} class="btn btn-primary btn-sm text-white"
            >Logout</button
        >
    </div>
</nav>
