<script>
  let fullName = '';
  let email = '';
  let phone = '';
  let citizenshipNumber = '';
  let dob = '';
  let gender = '';
  let address = '';
  let citizenshipFront;
  let citizenshipBack;
  let selfie;

  function handleSubmit() {
    console.log({
      fullName,
      email,
      phone,
      citizenshipNumber,
      dob,
      gender,
      address,
      citizenshipFront,
      citizenshipBack,
      selfie
    });
    alert('Form submitted successfully. We will verify your documents shortly.');
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100 p-6">
  <div class="w-full max-w-4xl bg-white rounded-xl shadow-lg p-10">
    <div class="text-center mb-8">
      <h1 class="text-3xl font-bold text-black">Digital Voter Registration</h1>
      <p class="text-gray-600 text-sm mt-2">Secure and verified government records.</p>
    </div>

    <form on:submit|preventDefault={handleSubmit} class="space-y-8">
      <div class="grid md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Full Name</label>
          <input type="text" bind:value={fullName} required placeholder="Ram Bahadur"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
          <input type="email" bind:value={email} required placeholder="you@example.com"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Phone </label>
          <input type="tel" bind:value={phone} required placeholder="9800000000"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Citizenship Number</label>
          <input type="text" bind:value={citizenshipNumber} required
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Date of Birth</label>
          <input type="date" bind:value={dob} required
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Gender</label>
          <select bind:value={gender} required
            class="w-full p-3 border border-gray-300 rounded-lg bg-white text-gray-700 focus:outline-none focus:ring-2 focus:ring-black">
            <option value="" disabled selected>Select Gender</option>
            <option value="Male">Male</option>
            <option value="Female">Female</option>
            <option value="Other">Other</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Permanent Address</label>
          <input type="text" bind:value={address} required placeholder="City, District"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
      </div>

      <div class="grid md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Citizenship Front</label>
          <input type="file" accept="image/*" required on:change={(e) => citizenshipFront = e.target.files[0]}
            class="block w-full text-base text-gray-700 border border-gray-300 rounded-lg cursor-pointer bg-white p-2 h-10 focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Citizenship Back</label>
          <input type="file" accept="image/*" required on:change={(e) => citizenshipBack = e.target.files[0]}
            class="block w-full text-base text-gray-700 border border-gray-300 rounded-lg cursor-pointer bg-white p-2 h-10 focus:outline-none focus:ring-2 focus:ring-black" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Selfie with Citizenship</label>
          <input type="file" accept="image/*" required on:change={(e) => selfie = e.target.files[0]}
            class="block w-full text-base text-gray-700 border border-gray-300 rounded-lg cursor-pointer bg-white p-2 h-10 focus:outline-none focus:ring-2 focus:ring-black" />
        </div>

      </div>

      <div class="flex justify-center">
        <button type="submit"
          class="bg-black text-white text-lg font-medium py-3 px-10 rounded-full hover:bg-gray-800 transition duration-300">
          Submit Registration
        </button>
      </div>

    </form>
  </div>
</div>
