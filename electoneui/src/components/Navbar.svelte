<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation"; // If using SvelteKit, remove if not

    let loggedIn = false;
    let mobileMenuOpen = false;
    let dropdownOpen = false;
    let profilePic = "";

    function toggleMenu() {
        mobileMenuOpen = !mobileMenuOpen;
    }

    function toggleDropdown() {
        dropdownOpen = !dropdownOpen;
    }

    function logout() {
        localStorage.removeItem("user_token");
        goto("/user/login"); // or window.location.href if not using SvelteKit
    }

    onMount(async () => {
        const token = localStorage.getItem("user_token");
        loggedIn = !!token;

        if (loggedIn) {
            const res = await fetch("http://localhost:3000/user/0", {
                headers: { Authorization: `Bearer ${token}` },
            });
            const data = await res.json();
            const photoPath = data.photos?.split(",")[3];
            profilePic = photoPath;
        }
    });
</script>

<nav
    class="flex items-center justify-between z-50 p-4 lg:px-8"
    aria-label="Global"
>
    <!-- Logo -->
    <div class="flex lg:flex-1">
        <a href={loggedIn ? "/user/dashboard" : "/"} class="-m-1.5 p-1.5">
            <span class="sr-only">ElectOne</span>
            <img class="h-8 w-auto" src="/favicon.png" />
        </a>
    </div>

    <!-- Desktop Nav Links -->
    <div class="hidden lg:flex lg:gap-x-12">
        {#if loggedIn}
            <!-- Could Be Anything Here? Deleted The Elections Nav As It Was Shifted -->
        {:else}
            <a href="/status" class="text-sm font-medium text-black">Status</a>
            <a href="/docs/about" class="text-sm font-medium text-black">About</a>
        {/if}
    </div>

    <!-- Desktop Profile + Dropdown -->
    <div class="hidden lg:flex lg:flex-1 lg:justify-end relative">
        {#if loggedIn}
            <button on:click={toggleDropdown}>
                <img
                    src={profilePic || "https://placehold.co/32x32"}
                    alt="Profile"
                    class="w-12 h-12 rounded-full object-cover"
                />
            </button>
            {#if dropdownOpen}
                <div
                    class="absolute right-0 mt-2 bg-white rounded shadow p-2 w-36 z-50 space-y-1"
                >
                    <a
                        href="/user/dashboard/profile"
                        class="block text-sm hover:underline">Profile</a
                    >
                    <button
                        on:click={logout}
                        class="text-sm text-red-500 hover:underline"
                        >Log Out</button
                    >
                </div>
            {/if}
        {:else}
            <div class="flex items-center gap-4">
                <a href="/user/signup" class="text-sm font-semibold text-black"
                    >Signup</a
                >
                <a
                    href="/user/login"
                    class="text-sm font-semibold btn btn-primary btn-sm rounded-lg"
                >
                    Login <span aria-hidden="true">&rarr;</span>
                </a>
            </div>
        {/if}
    </div>

    <!-- Mobile Hamburger -->
    <button class="block lg:hidden md:hidden" on:click={toggleMenu}>
        <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            viewBox="0 0 24 24"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4 6h16M4 12h16M4 18h16"
            />
        </svg>
    </button>
</nav>

<!-- Mobile Menu -->
{#if mobileMenuOpen}
    <div class="lg:hidden px-4 pb-4 space-y-3">
        {#if loggedIn}
            <a href="/user/dashboard" class="block text-sm text-black"
                >Dashboard</a
            >
            <a href="/user/dashboard/profile" class="block text-sm text-black"
                >Profile</a
            >
            <button on:click={logout} class="block text-sm text-red-500"
                >Log Out</button
            >
        {:else}
            <a href="/status" class="block text-sm text-black">Status</a>
            <a href="#about" class="block text-sm text-black">About</a>
            <a href="/user/signup" class="block text-sm text-black mt-2"
                >Signup</a
            >
            <a
                href="/user/login"
                class="block text-sm btn btn-primary btn-sm rounded-lg mt-2 w-full text-center"
            >
                Login &rarr;
            </a>
        {/if}
    </div>
{/if}
