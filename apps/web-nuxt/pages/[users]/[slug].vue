<script lang="ts" setup>
// const route = useRoute()
const tags = ['tag1', 'tag2']
const source = `# Turborepo starter

[![codecov](https://codecov.io/github/nomorechokedboy/pizza/branch/main/graph/badge.svg?token=8USH85MRVT)](https://codecov.io/github/nomorechokedboy/pizza)
[![CI workflows](https://github.com/nomorechokedboy/pizza/actions/workflows/ci.yml/badge.svg)](https://github.com/nomorechokedboy/pizza/actions/workflows/ci.yml)
[![golangci-lint](https://github.com/nomorechokedboy/pizza/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/nomorechokedboy/pizza/actions/workflows/golangci-lint.yml)

This is an official pnpm starter turborepo. The monorepo contains frontend, cms written in [VueJS](https://vuejs.org) and api written in [Go](https://go.dev/)

## Prerequisite

Because of the mix with typescript and golang, you will need to install go dependencies if you wish to run the api code.

And of course, please install [NodeJs](https://nodejs.dev/) and [Golang](https://go.dev/) locally for development.

### Set GOPATH to use go binary

Running this command will set your environtment path with $GOPATH/bin so you can use golang binary commands

\`\`\`sh
echo 'export PATH=$(go env GOPATH)/bin:$PATH' >> ~/.zshrc # (or .bashrc)
\`\`\`

### Swaggo

Install swaggo with the below installation command [from fiber swagger repo](https://github.com/gofiber/swagger)

\`\`\`sh
go get -u github.com/swaggo/swag/cmd/swag
# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest

swag -v
\`\`\`

### Golanglint-ci

To install Golanglint-ci, run the command from golanglint-ci installation guide. If you wish to install with other methods, you can checkout [the documentation](https://golangci-lint.run/usage/install/)

\`\`\`sh
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

golangci-lint --version
\`\`\`

After this, you are good to contribute to the golang repository

## What's inside?

This turborepo uses [pnpm](https://pnpm.io) as a package manager. It includes the following packages/apps:

### Dev container

This monorepo use dev containers for development, you could choose between local environment and dev containers environment for developing the project.

Just install dev containers and run it in the container

_Note_: dev containers is pretty new with me so the setup is not at its best so you might have some trouble with dev containers.

### Apps and Packages

- \`admin-nuxt\`: a [Nuxt.js](https://nuxt.com) app
- \`web-nuxt\`: a [Nuxt.js](https://nuxt.com) app
- \`web\`: another [Next.js](https://nextjs.org) app
- \`ui\`: a stub React component library shared by \`web\` applications
- \`ui\`: a stub Vue component library shared by both \`admin-nuxt\` and \`web-nuxt\` applications
- \`eslint-config-custom\`: \`eslint\` configurations (includes \`eslint-config-next\` and \`eslint-config-prettier\`)
- \`tsconfig\`: \`tsconfig.json\`s used throughout the monorepo

Each package/app is 100% [TypeScript](https://www.typescriptlang.org/) and [Golang](https://go.dev/)

### Utilities

This turborepo has some additional tools already setup for you:

- [TypeScript](https://www.typescriptlang.org/) for static type checking
- [ESLint](https://eslint.org/) for code linting
- [Prettier](https://prettier.io) for code formatting

### Build

To build all apps and packages, run the following command:

\`\`\`
cd pizza
pnpm run build
\`\`\`

### Develop

To develop all apps and packages, run the following command:

\`\`\`
cd pizza
pnpm run dev
\`\`\`

### Remote Caching

Turborepo can use a technique known as [Remote Caching](https://turborepo.org/docs/core-concepts/remote-caching) to share cache artifacts across machines, enabling you to share build caches with your team and CI/CD pipelines.

By default, Turborepo will cache locally. To enable Remote Caching you will need an account with Vercel. If you don't have an account you can [create one](https://vercel.com/signup), then enter the following commands:

\`\`\`
cd pizza
pnpm dlx turbo login
\`\`\`

This will authenticate the Turborepo CLI with your [Vercel account](https://vercel.com/docs/concepts/personal-accounts/overview).

Next, you can link your Turborepo to your Remote Cache by running the following command from the root of your turborepo:

\`\`\`
pnpm dlx turbo link
\`\`\`

## Useful Links

Learn more about the power of Turborepo:

- [Pipelines](https://turborepo.org/docs/core-concepts/pipelines)
- [Caching](https://turborepo.org/docs/core-concepts/caching)
- [Remote Caching](https://turborepo.org/docs/core-concepts/remote-caching)
- [Scoped Tasks](https://turborepo.org/docs/core-concepts/scopes)
- [Configuration Options](https://turborepo.org/docs/reference/configuration)
- [CLI Usage](https://turborepo.org/docs/reference/command-line-reference)
`
</script>

<template>
	<SidebarLayout class="md:p-4 lg:max-w-7xl">
		<template #sidebar>
			<LeftSidebar class="lg:max-w-[64px]" />
		</template>
		<SidebarLayout :reverse="true">
			<template #sidebar>
				<LeftSidebar
					class="hidden lg:block lg:max-w-xs"
				/>
			</template>
			<div class="w-full max-w-[832px]">
				<article
					class="bg-white border border-neutral-200"
				>
					<nuxt-img
						sizes="sm:100vw md:680px, 806px"
						src="https://res.cloudinary.com/practicaldev/image/fetch/s--h0iBOiYw--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qxfnu4dptk7iazjv6mjg.png"
					/>
					<section
						class="flex flex-col gap-3 p-4"
					>
						<NoClue
							description="Lmao"
							src="https://api.dicebear.com/6.x/icons/svg?seed=CheeseRaa"
							title="owner.name"
						/>
						<h1
							class="text-3xl font-bold text-neutral-900"
						>
							Learn How to Use
							Artifactory as a Docker
							Registry
						</h1>
						<Tags :tags="tags" />
						<main class="max-w-full">
							<VueMarkdown
								:source="source"
							/>
						</main>
					</section>
					<CommentSection />
				</article>
			</div>
		</SidebarLayout>
	</SidebarLayout>
</template>
