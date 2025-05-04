<script>
    import { onMount } from "svelte";

    let loggedIn = false;

    // Check localStorage on mount
    onMount(() => {
        loggedIn = !!localStorage.getItem("token");
    });

    let mobileMenuOpen = false;

    function toggleMenu() {
        mobileMenuOpen = !mobileMenuOpen;
    }
</script>

<nav class="flex items-center justify-between p-4 bg-gray-900 text-white">
    <!-- Logo -->
    <div class="text-xl font-bold">E</div>

    <!-- Desktop Menu -->
    <div class="hidden md:flex gap-4 items-center">
        {#if loggedIn}
            <a href="/user/dashboard/" class="hover:text-gray-300">News</a>
            <a href="/user/dashboard/election" class="hover:text-gray-300"
                >Elections</a
            >
            <a href="/user/profile/">
                <img
                    src="https://placehold.co/32x32"
                    alt="Profile"
                    class="w-8 h-8 rounded-full"
                />
            </a>
        {:else}
            <a href="/user/login" class="hover:text-gray-300">Login</a>
            <a href="/user/signup" class="hover:text-gray-300">Register</a>
        {/if}
    </div>

    <!-- Hamburger -->
    <button class="md:hidden" on:click={toggleMenu}>
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

    <!-- Mobile Menu -->
    {#if mobileMenuOpen}
        <div
            class="absolute top-16 right-4 bg-gray-800 p-4 rounded shadow-md flex flex-col gap-2 md:hidden"
        >
            {#if loggedIn}
                <a href="/news" class="hover:text-gray-300">News</a>
                <a href="/elections" class="hover:text-gray-300">Elections</a>
                <a href="/profile" class="flex items-center gap-2">
                    <img
                        src="https://placehold.co/32x32"
                        alt="Profile"
                        class="w-8 h-8 rounded-full"
                    />
                    <span>Profile</span>
                </a>
            {:else}
                <a href="/login" class="hover:text-gray-300">Login</a>
                <a href="/register" class="hover:text-gray-300">Register</a>
            {/if}
        </div>
    {/if}
</nav>

<style>
    nav {
        position: relative;
    }
</style>
