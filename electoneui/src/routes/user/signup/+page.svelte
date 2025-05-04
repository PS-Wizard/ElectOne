<script>
    import { onMount } from "svelte";
    let citizenship_id = "";
    let voter_card_id = "";
    let password = "";
    let photos = ""; // Will store comma-separated filenames

    const handleSubmit = async () => {
      const appealData = {
        citizenship_id: citizenship_id,
        voter_card_id: voter_card_id,
        password: password,
        photos: photos,
      };

      const authToken =
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbklEIjoyLCJlbWFpbCI6ImFkbWluMkBleGFtcGxlLmNvbSIsImV4cCI6MTc0NjQ0NTcyNywicm9sZSI6ImFkbWluIn0.9msec4uS7vVWaMIpC47jV1ne-iO_kvvC-bWJRt4UjX8"; // Replace with actual token handling

      try {
        const response = await fetch("http://localhost:3000/appeal", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${authToken}`,
          },
          body: JSON.stringify(appealData),
        });

        if (response.ok) {
          alert("Appeal submitted successfully!");
          // Optionally redirect the user
          // window.location.href = "/appeal-submitted";
        } else {
          const error = await response.text();
          alert(`Error submitting appeal: ${error}`);
        }
      } catch (err) {
        console.error("Network error:", err);
        alert("Network error. Please try again later.");
      }
    };
  </script>

  <svelte:head>
    <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
  </svelte:head>

  <div class="min-h-screen flex items-center justify-center bg-gray-100 p-6">
    <div class="w-full max-w-md bg-white rounded-xl shadow-lg p-10">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-black">Create Appeal</h1>
        <p class="text-gray-600 text-sm mt-2">Submit your appeal details.</p>
      </div>

      <form on:submit|preventDefault={handleSubmit} class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1"
            >Citizenship ID</label
          >
          <input
            type="text"
            bind:value={citizenship_id}
            required
            placeholder="123-456"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1"
            >Voter Card ID</label
          >
          <input
            type="text"
            bind:value={voter_card_id}
            required
            placeholder="VC-001"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1"
            >Password</label
          >
          <input
            type="password"
            bind:value={password}
            required
            placeholder="securepassword123"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1"
            >Photo Filenames (comma-separated)</label
          >
          <input
            type="text"
            bind:value={photos}
            required
            placeholder="photo1.jpg,photo2.png,photo3.jpeg"
            class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-black"
          />
        </div>

        <div class="flex justify-center">
          <button
            type="submit"
            class="bg-black text-white text-lg font-medium py-3 px-10 rounded-full hover:bg-gray-800 transition duration-300"
          >
            Submit Appeal
          </button>
        </div>
      </form>
    </div>
  </div>