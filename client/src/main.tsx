import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { AuthProvider } from './context/AuthProvider';
import { LoginForm } from './components/Forms/login-form.tsx';
import { SignupForm } from './components/Forms/signup-form.tsx';
import { PasswordResetForm } from './components/Forms/passwordreset-form.tsx';
import { ReleasesPage } from './components/pages/releases-page.tsx';
import { AboutPage } from './components/pages/about-page.tsx';
import { GalleryPage } from './components/pages/gallery-page.tsx';
import { WebstorePage } from './components/pages/webstore-page.tsx';
import { NewsPage } from './components/pages/news-page.tsx';
import { NotFoundPage } from './components/pages/notfound-page.tsx';

const router = createBrowserRouter([
  { 
    path: "/", 
    element: <App />,
    children: [
      { path: "/login", element: <LoginForm className="w-100" /> },
      { path: "/signup", element: <SignupForm className="w-100" /> },
      { path: "/reset-password", element: <PasswordResetForm className="w-100" /> },
      { path: "/about", element: <AboutPage /> },
      { path: "/gallery", element: <GalleryPage /> },
      { path: "/releases", element: <ReleasesPage /> },
      { path: "/webstore", element: <WebstorePage /> },
      { path: "/news", element: <NewsPage /> },
    ]
  },
  { path: "*", element: <NotFoundPage /> },
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <div className="bg-linear-to-r from-gray-700 via-black to-gray-700 text-white w-screen h-screen absolute top-0 left-0 flex items-center justify-center">
        <RouterProvider router={router} />
      </div>
    </AuthProvider>
  </StrictMode>,
)
