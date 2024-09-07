<script>
    export let error;
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100 text-center">
    <div class="max-w-sm p-6 bg-white rounded-lg shadow-lg">
        <h1 class="text-4xl font-bold text-red-600">400</h1>
        <p class="mt-4 text-lg text-gray-700">¡Ups! Algo salió mal.</p>
        <p class="mt-4 text-gray-500">{error}</p>
        <a href="/" class="mt-8 inline-block px-4 py-2 bg-blue-500 text-white rounded-lg">Volver al inicio</a>
    </div>
</div>
