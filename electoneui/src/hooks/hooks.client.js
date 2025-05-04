export async function handleFetch({ request, fetch, options }) {
	const token = localStorage.getItem('token');

	options.headers = {
		...options.headers,
		Authorization: token ? `Bearer ${token}` : '',
	};

	return fetch(request, options);
}

