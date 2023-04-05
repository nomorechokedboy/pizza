<script lang="ts" setup>
import { ActionIcon, Button } from 'ui-vue'
import ChevronIcon from '~icons/mdi/chevron-up-down'
import { CommentProps } from './Comment.vue'

const source = `
## Getting Started

First, run the development server:

\`\`\`bash
yarn dev
\`\`\`

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying \`pages/index.js\`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited in \`pages/api/hello.js\`.

The \`pages/api\` directory is mapped to \`/api/*\`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_source=github.com&utm_medium=referral&utm_campaign=turborepo-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
`

const comments: CommentProps[] = [
	{
		content: source,
		createdAt: 'Mar 28',
		updated: true,
		like: 5,
		user: {
			name: 'Duong Le',
			avatarUrl: 'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif'
		},
		replies: []
	},
	{
		content: source,
		createdAt: 'Mar 28',
		updated: false,
		like: 15,
		user: {
			name: 'Duong Le',
			avatarUrl: 'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif'
		},
		replies: [
			{
				content: source,
				createdAt: 'Mar 28',
				updated: false,
				like: 15,
				user: {
					name: 'Duong Le',
					avatarUrl: 'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif'
				},
				replies: []
			},
			{
				content: source,
				createdAt: 'Mar 28',
				updated: false,
				like: 15,
				user: {
					name: 'Duong Le',
					avatarUrl: 'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif'
				},
				replies: [
					{
						content: source,
						createdAt: 'Mar 28',
						updated: false,
						like: 15,
						user: {
							name: 'Duong Le',
							avatarUrl: 'https://res.cloudinary.com/practicaldev/image/fetch/s--5VEqFAA8--/c_fill,f_auto,fl_progressive,h_90,q_66,w_90/https://dev-to-uploads.s3.amazonaws.com/uploads/user/profile_image/909049/9a19683f-1e9f-4933-bdba-e7ea2fe5e71c.gif'
						},
						replies: []
					}
				]
			}
		]
	}
]
</script>

<template>
	<section class="p-4 border-t border-t-neutral-200 flex flex-col gap-6">
		<header class="flex items-center justify-between">
			<div class="flex items-center">
				<h2 class="text-xl text-neutral-800 font-bold">
					Top comments (0)
				</h2>
				<ActionIcon
					color="indigo"
					variant="subtle"
					size="xl"
				>
					<ChevronIcon
						class="text-neutral-800"
						width="24"
					/>
				</ActionIcon>
			</div>
			<Button>Subscribe</Button>
		</header>
		<Comment
			v-for="(comment, i) in comments"
			v-bind="comment"
			:key="i"
		/>
		<nav
			class="flex items-center justify-center text-neutral-500 text-sm"
		>
			<p>Code of conduct</p>
			<div class="py-2 px-3">
				<div
					class="h-1 w-1 rounded-full bg-[rgb(189,189,189)]"
				/>
			</div>
			<p>Report abuse</p>
		</nav>
	</section>
</template>
