<script>
    import Footer from "../../../components/footer.svelte";
    import Navbar from "../../../components/Navbar.svelte";

    let lastUpdated = new Date().toLocaleDateString();
</script>

<Navbar />
<div
    class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
    aria-hidden="true"
>
    <div
        class="relative left-[calc(80%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
        style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
    ></div>
</div>
<div
    class="max-w-3xl mx-auto my-12 p-6 rounded-2xl font-sans text-black"
>
    <h1 class="text-3xl font-bold underline mb-4">Cookie Policy</h1>
    <p class="text-sm text-gray-500 mb-8">Last updated: {lastUpdated}</p>

    <h2 class="text-2xl font-semibold mb-2">What Are Cookies?</h2>
    <p class="mb-6">
        Cookies are small text files stored on your device when you visit a
        website. They help us improve your experience by remembering preferences
        and user behavior.
    </p>

    <h2 class="text-2xl font-semibold mb-2">How We Use Cookies</h2>
    <p class="mb-2">Our voting website uses cookies to:</p>
    <ul class="list-disc pl-6 mb-6">
        <li>Keep you logged in during your session</li>
        <li>Ensure secure voting processes</li>
        <li>Analyze website traffic and usage patterns</li>
    </ul>

    <h2 class="text-2xl font-semibold mb-2">Types of Cookies We Use</h2>
    <ul class="list-disc pl-6 mb-6">
        <li>
            <strong>Essential Cookies:</strong> Required for the site to function
            correctly
        </li>
    </ul>

    <h2 class="text-2xl font-semibold mb-2">Managing Cookies</h2>
    <p class="mb-6">
        You can control or delete cookies through your browser settings.
        Disabling cookies may affect your experience on our site.
    </p>

    <h2 class="text-2xl font-semibold mb-2">Third-Party Cookies</h2>
    <p class="mb-6">
        Some features may use third-party services that also use cookies. These
        are subject to the respective providers' privacy policies.
    </p>

    <h2 class="text-2xl font-semibold mb-2">Contact Us</h2>
    <p>
        If you have questions about our use of cookies, please contact us at
        <a
            href="mailto:privacy@example.com"
            class="text-blue-500 hover:underline">privacy@example.com</a
        >.
    </p>
</div>

<Footer />
