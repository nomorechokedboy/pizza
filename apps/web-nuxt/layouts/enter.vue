<template>
	<NuxtLayout name="default">
		<main
			class="flex flex-col flex-1 lg:flex-row gap-8 bg-gray-300 lg:p-8"
		>
			<div class="flex flex-col flex-1 justify-center">
				<div
					class="flex-1 flex flex-col items-center bg-white p-6 gap-6"
				>
					<header class="w-full">
						<h1
							class="text-center font-bold text-2xl"
						>
							Welcome to Copium
							Community
						</h1>
						<p class="text-center">
							Copium Community is a
							Community of 1,026,041
							amazing copers
						</p>
					</header>
					<slot />
				</div>
			</div>
		</main>
	</NuxtLayout>
</template>
