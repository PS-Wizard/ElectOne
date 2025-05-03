<script>
    const links = [
        { label: "Dashboard", path: "/admin/dashboard" },
        { label: "Appeals", path: "/admin/appeals" },
        { label: "Candidates", path: "/admin/candidates" },
        { label: "Citizenships", path: "/admin/citizenships" },
        { label: "Elections", path: "/admin/electons" },
        { label: "Users", path: "/admin/users" },
        { label: "Voter Cards", path: "/admin/votercards" },
    ];
</script>

<div
    class="max-w-2xl mx-auto mt-10 p-6 bg-base-200 rounded-xl shadow space-y-4"
>
    <h1 class="text-2xl font-bold mb-4">Admin Dashboard</h1>

    <div class="grid gap-3">
        {#each links as link}
            <a
                href={link.path}
                class="btn btn-outline btn-primary w-full justify-start"
            >
                {link.label}
            </a>
        {/each}
    </div>
</div>
