<script>
    import Footer from "../../../components/footer.svelte";
    import Navbar from "../../../components/Navbar.svelte";
    // import AdminNavbar from "../../../components/AdminNavbar.svelte";
</script>

<!-- Full Page Layout -->

<Navbar />
<div
    class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80"
    aria-hidden="true"
>
    <div
        class="relative left-[calc(50%-11rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
        style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
    ></div>
</div>
<div class="pb-12 bg-transparent min-h-screen flex flex-col">
    <!-- Top Navbar with Back Arrow and Title -->
    <div class="flex items-center justify-between px-6 py-4 bg-transparent">
        <!-- Back Arrow with no navigation -->
        <button
            on:click={() => window.history.back()}
            class="text-3xl text-black font-bold cursor-pointer">←</button
        >
        <!-- Empty space for balancing the layout -->
    </div>

    <!-- Privacy Policy Content -->
    <div class="w-full text-gray-700 py-10 px-6">
        <h1 class="text-6xl font-bold text-gray-900 text-center mb-6">
            Privacy Policy
        </h1>
        <p class="text-sm text-gray-600 mb-6 text-center">
            At Online Voting System Nepal, your privacy is our priority. This
            policy explains how we collect, use, and protect your personal
            information.
        </p>

        <!-- Policy Sections -->
        <div class="space-y-10 text-base text-left max-w-5xl mx-auto">
            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    1. Information We Collect
                </h2>
                <ul class="text-base list-disc ml-6 space-y-1">
                    <li>Full Name</li>
                    <li>Voter ID / Citizenship Number</li>
                    <li>Login Information</li>
                </ul>
            </section>

            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    2. How We Use Your Information
                </h2>
                <ul class="list-disc ml-6 space-y-1">
                    <li>Verify voter identity</li>
                    <li>Provide secure login access</li>
                    <li>Prevent duplicate or fraudulent voting</li>
                    <li>Generate anonymized election statistics</li>
                    <li>Improve system security and user experience</li>
                </ul>
                <p class="mt-2">
                    We never sell or share your data with third parties without
                    consent, except as required by law.
                </p>
            </section>

            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    3. Data Protection Measures
                </h2>
                <ul class="list-disc ml-6 space-y-1">
                    <li>All personal data is encrypted end-to-end.</li>
                    <li>
                        Votes are stored anonymously and cannot be traced back
                        to you.
                    </li>
                    <li>
                        Regular audits and real-time security monitoring are in
                        place.
                    </li>
                    <li>
                        Multi-factor authentication and OTP verification are
                        mandatory.
                    </li>
                </ul>
            </section>

            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    4. Your Rights
                </h2>
                <ul class="list-disc ml-6 space-y-1">
                    <li>Access and review your registered data</li>
                    <li>
                        Request correction of incorrect data <span
                            class="text-red-600">(upon verification)</span
                        >
                    </li>
                    <li>
                        Request account deletion <span class="text-red-600"
                            >(before voting period)</span
                        >
                    </li>
                </ul>
                <p class="mt-2">
                    For concerns, email us at <span
                        class="underline text-blue-600"
                        >privacy@electonesystemnp.org</span
                    >
                </p>
            </section>

            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    5. Cookies & Tracking
                </h2>
                <p>
                    We use minimal tracking for system performance and error
                    logging. No third-party ads or trackers are used. We do not
                    store cookies for marketing purposes.
                </p>
            </section>

            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    6. Policy Updates
                </h2>
                <p>
                    Our privacy policy may be updated periodically. All changes
                    will be communicated through app notifications and official
                    OTP.
                </p>
            </section>

            <!-- ✅ New Section 7: Data Retention Policy -->
            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    7. Data Retention Policy
                </h2>
                <p>
                    All personal and voting-related data is retained only as
                    long as necessary to fulfill legal and audit requirements.
                    After the conclusion of the election cycle:
                </p>
                <ul class="list-disc ml-6 space-y-1 mt-2">
                    <li>
                        Personally identifiable information is securely archived
                        or deleted based on retention schedules.
                    </li>
                    <li>
                        Vote records remain anonymized and untraceable to ensure
                        election integrity without compromising user privacy.
                    </li>
                    <li>
                        All data deletion requests are processed in compliance
                        with Nepalese data protection laws.
                    </li>
                </ul>
            </section>

            <!-- ✅ New Section 8: Third-Party Services and Compliance -->
            <section>
                <h2 class="text-3xl font-semibold text-gray-800 mb-2">
                    8. Third-Party Services and Compliance
                </h2>
                <p>
                    Our system may integrate with third-party tools for OTP
                    verification, identity checks, or cloud infrastructure. We
                    ensure:
                </p>
                <ul class="list-disc ml-6 space-y-1 mt-2">
                    <li>
                        All third-party services comply with Nepali data
                        protection laws.
                    </li>
                    <li>
                        No personal data is shared without purpose-specific user
                        consent.
                    </li>
                    <li>
                        We conduct regular security and compliance reviews of
                        all integrated services.
                    </li>
                </ul>
            </section>
        </div>
    </div>
    <div
        class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
        aria-hidden="true"
    >
        <div
            class="relative left-[calc(50%+3rem)] aspect-1155/678 w-[36.125rem] -translate-x-1/2 bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
            style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)"
        ></div>
    </div>
</div>


<Footer />
