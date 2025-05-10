<script>
    import { onMount } from "svelte";

    let loggedIn = false;

    onMount(() => {
        loggedIn = !!localStorage.getItem("user_token");
    });

    let mobileMenuOpen = false;

    function toggleMenu() {
        mobileMenuOpen = !mobileMenuOpen;
    }
</script>

<nav class="flex items-center justify-between p-4 lg:px-8" aria-label="Global">
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

    <!-- Desktop Nav -->
    <div class="hidden lg:flex lg:gap-x-12">
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
            <a href="#about" class="text-sm font-medium text-black">About</a>
        {/if}
    </div>

    <!-- Desktop Right -->
    <div class="hidden lg:flex lg:flex-1 lg:justify-end">
        {#if loggedIn}
            <a href="/user/dashboard/profile">
                <img
                    src="https://placehold.co/32x32"
                    alt="Profile"
                    class="w-8 h-8 rounded-full"
                />
            </a>
        {:else}
            <div class="flex items-center gap-4">
                <a href="/user/signup" class="text-sm font-semibold text-black"
                    >Signup</a
                >
                <a
                    href="/user/login"
                    class="text-sm font-semibold btn btn-primary btn-sm rounded-lg"
                    >Login <span aria-hidden="true">&rarr;</span></a
                >
            </div>
        {/if}
    </div>

    <!-- Hamburger (Mobile) -->
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

<!-- Mobile Dropdown -->
{#if mobileMenuOpen}
    <div class="lg:hidden px-4 pb-4 space-y-3">
        {#if loggedIn}
            <a href="/user/dashboard" class="block text-sm text-black">News</a>
            <a href="/user/dashboard/election" class="block text-sm text-black"
                >Elections</a
            >
            <a href="/user/profile" class="flex items-center gap-2 mt-2">
                <img
                    src="https://placehold.co/32x32"
                    alt="Profile"
                    class="w-8 h-8 rounded-full"
                />
                <span class="text-sm text-black">Profile</span>
            </a>
        {:else}
            <a href="#features" class="block text-sm text-black">Features</a>
            <a href="#about" class="block text-sm text-black">About</a>
            <a href="/user/signup" class="block text-sm text-black mt-2"
                >Signup</a
            >
            <a
                href="/user/login"
                class="block text-sm btn btn-primary btn-sm rounded-lg mt-2 w-full text-center"
                >Login &rarr;</a
            >
        {/if}
    </div>
{/if}
