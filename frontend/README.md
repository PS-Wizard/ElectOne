# Frontend Zone:

- Create Components In Their Own Seperate File
- Will be merged to the main branch later


# Dependency:
`tailwindcss` is the only dependency

```html
~

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTMX + Tailwind</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-100 text-gray-900 flex items-center justify-center min-h-screen">
    <div class="p-6 bg-white shadow-lg rounded-lg text-center">
        <h1 class="text-2xl font-bold">HTMX + Tailwind Boilerplate</h1>
        <button 
            class="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            hx-get="/your-endpoint"
            hx-target="#response">
            Click Me
        </button>
        <div id="response" class="mt-4"></div>
    </div>
</body>
</html>

```
