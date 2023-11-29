import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import NavBar from '@/components/navigation/NavBar'
import { AppProvider } from '@/providers/Auth0Provider'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'React-Golang-Auth0-Template',
  description: 'This is a template for using React, Golang, and Auth0',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <AppProvider>
    <html lang="en">
      <body className={inter.className}>
        <NavBar />
        {children}
      </body>
    </html>
    </AppProvider>
  )
}
