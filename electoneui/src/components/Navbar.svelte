<script>
    import { onMount } from "svelte";

    let loggedIn = false;

    // Check localStorage on mount
    onMount(() => {
        loggedIn = !!localStorage.getItem("user_token");
    });

    let mobileMenuOpen = false;

    function toggleMenu() {
        mobileMenuOpen = !mobileMenuOpen;
    }
</script>

<header class="absolute inset-x-0 top-0 z-50">
    <nav
        class="flex items-center justify-between p-6 lg:px-8"
        aria-label="Global"
    >
        <div class="flex lg:flex-1">
            <a href={loggedIn ? "/user/dashboard" : "/"} class="-m-1.5 p-1.5">
                <span class="sr-only">ElectOne</span>
                <img
                    class="h-8 w-auto"
                    src="https://via.placeholder.com/100x50.png?text=E"
                    alt="E"
                />
            </a>
        </div>
        <div class="lg:flex lg:gap-x-12">
            {#if loggedIn}
                <a href="/user/dashboard" class="text-sm font-medium text-black"
                    >News</a
                >
                <a
                    href="/user/dashboard/election"
                    class="text-sm font-medium text-black">Elections</a
                >
            {:else}
                <a href="#features" class="text-sm font-medium text-black"
                    >Features</a
                >
                <a href="#about" class="text-sm font-medium text-black">About</a
                >
            {/if}
        </div>
        <div class="lg:flex lg:flex-1 lg:justify-end">
            {#if loggedIn}
                <a href="/user/profile/">
                    <img
                        src="https://placehold.co/32x32"
                        alt="Profile"
                        class="w-8 h-8 rounded-full"
                    />
                </a>
            {:else}
                <div class="flex justify-center items-center gap-10">
                    <a
                        href="/user/signup"
                        class="text-sm font-semibold text-black">Signup</a
                    >
                    <a
                        href="/user/login"
                        class="text-sm/6 font-semibold btn btn-primary btn-sm rounded-lg"
                        >Login<span aria-hidden="true">&rarr;</span></a
                    >
                </div>
            {/if}
        </div>
    </nav>

    <!-- Mobile Menu -->
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
</header>

<style>
    nav {
        position: relative;
    }
</style>
