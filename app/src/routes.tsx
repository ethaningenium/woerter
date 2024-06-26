import { Outlet, createBrowserRouter } from "react-router-dom";
import { Home } from "@/features/home";
import { CreateWordPage } from "@/features/create_card";
import { WordPage } from "@/features/word";
import { EditPage } from "@/features/edit_card";
import { Onboarding } from "@/features/onboarding";
import { AskPage } from "@/features/ask";
import { GoPlayPage, PlayPage } from "@/features/play";
import { ProfilePage } from "@/features/profile";
import { LoginPage } from "@/features/login";
import { AuthProvider, Header } from "./features/common";

export const router = createBrowserRouter([
  {
    path: "/app/login",
    element: <LoginPage />,
  },
  {
    path: "/",
    element: (
      <AuthProvider>
        <Header />
        <Outlet />
      </AuthProvider>
    ),
    children: [
      {
        path: "/app",
        element: <Home />,
      },
      {
        path: "/app/play",
        element: <PlayPage />,
      },
      {
        path: "/app/profile",
        element: <ProfilePage />,
      },
      {
        path: "/app/goplay",
        element: <GoPlayPage />,
      },
      {
        path: "/app/create/:lang",
        element: <CreateWordPage />,
      },
      {
        path: "/app/dic/:id",
        element: <WordPage />,
      },
      {
        path: "/app/edit/:id",
        element: <EditPage />,
      },
      {
        path: "/app/onboard",
        element: <Onboarding />,
      },
      {
        path: "/app/ask/:lang",
        element: <AskPage />,
      },
    ],
  },
]);
