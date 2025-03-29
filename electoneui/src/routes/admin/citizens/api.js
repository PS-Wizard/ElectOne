const BASE_URL = "https://localhost:3000/api/secure";

export async function GetCitizens(page = 0) {
    const res = await fetch(`${BASE_URL}/citizensPaginated/${page}`, {
        method: "GET",
        credentials: "include"
    });
    
    if (!res.ok) {
        console.error('Failed to fetch citizens:', res.status);
        return [];  // Return empty array if the fetch fails
    }

    const data = await res.json();
    console.log("Fetched citizens:", data);  // Log the data for debugging
    return data;
}

export async function DeleteCitizen(id) {
    const res = await fetch(`${BASE_URL}/citizens/${id}`, {
        method: "DELETE",
        credentials: "include"
    });
    return res.json();
}

export async function UpdateCitizen(id, citizen) {
    const res = await fetch(`${BASE_URL}/citizens/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(citizen)
    });
    return res.json();
}

export async function CreateCitizen(citizen) {
    const res = await fetch(`${BASE_URL}/citizens`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(citizen)
    });
    return res.json();
}
