<script>
    import Footer from "../../../components/footer.svelte";
    import Navbar from "../../../components/Navbar.svelte";
    // import AdminNavbar from "../../../components/AdminNavbar.svelte";
</script>

<!-- Full Page Layout -->

<div class="bg-gray-100 min-h-screen flex flex-col">
    <Navbar />
    <!-- Top Navbar with Back Arrow and Title -->
    <div class="flex items-center justify-between px-6 py-4 bg-white shadow-md">
        <!-- Back Arrow with no navigation -->
        <button
            on:click={() => window.history.back()}
            class="text-xl text-black font-bold cursor-pointer">←</button
        >
        <h1 class="text-3xl font-bold text-blue-600 text-center w-full">
            Elect One
        </h1>
        <div class="w-10"></div>
        <!-- Empty space for balancing the layout -->
    </div>

    <!-- Privacy Policy Content -->
    <div class="w-full bg-white text-gray-700 py-10 px-6">
        <h1 class="text-3xl font-bold text-gray-900 text-center mb-6">
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
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
                    1. Information We Collect
                </h2>
                <ul class="list-disc ml-6 space-y-1">
                    <li>Full Name</li>
                    <li>Voter ID / Citizenship Number</li>
                    <li>Login Information</li>
                </ul>
            </section>

            <section>
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
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
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
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
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
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
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
                    5. Cookies & Tracking
                </h2>
                <p>
                    We use minimal tracking for system performance and error
                    logging. No third-party ads or trackers are used. We do not
                    store cookies for marketing purposes.
                </p>
            </section>

            <section>
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
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
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
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
                <h2 class="text-xl font-semibold text-gray-800 mb-2">
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
</div>

<Footer />
