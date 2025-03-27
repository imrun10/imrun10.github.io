# Personal Portfolio Website

This is my personal portfolio website built using **Next.js** and **Supabase**. It showcases my projects, skills, and experiences, providing a professional online presence.

## Features
- **Modern UI**: Responsive and elegant design with a clean layout.
- **Projects Showcase**: Display portfolio projects with descriptions and links.
- **Contact Form**: Users can reach out through a contact form integrated with Supabase.
- **Blog (Optional)**: Share insights, tutorials, or updates.
- **Authentication (Optional)**: Secure login for admin functionalities.
- **Database Storage**: Supabase for handling project data, user messages, and authentication.

## Tech Stack
- **Framework**: Next.js (React-based framework for SSR and SSG)
- **Backend**: Supabase (PostgreSQL, authentication, and real-time features)
- **UI Library**: Tailwind CSS (for styling)
- **Deployment**: Vercel / Netlify

## Installation & Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/imrun10/imrun10.github.io.git
   cd imrun10.github.io
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
   or
   ```sh
   yarn install
   ```
3. Set up environment variables:
   Create a `.env.local` file and add your Supabase credentials:
   ```env
   NEXT_PUBLIC_SUPABASE_URL=your_supabase_url
   NEXT_PUBLIC_SUPABASE_ANON_KEY=your_supabase_anon_key
   ```
4. Run the development server:
   ```sh
   npm run dev
   ```
   or
   ```sh
   yarn dev
   ```
5. Open `http://localhost:3000/` in your browser.

## Deployment
You can deploy the website to Vercel, Netlify, or any hosting provider that supports Next.js.

For Vercel:
1. Install Vercel CLI:
   ```sh
   npm install -g vercel
   ```
2. Deploy:
   ```sh
   vercel
   ```

## Future Improvements
- Add animations and transitions for better UI experience.
- Implement a CMS for blog management.
- Integrate more social media links and analytics.

## License
This project is open-source and available under the MIT License.

