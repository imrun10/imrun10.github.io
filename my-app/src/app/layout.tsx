"use client"


import React from 'react';
import './globals.css';

import { Inter } from 'next/font/google';
import { CssBaseline } from '@mui/material';

import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers';
import Header from '../components/Header';

const inter = Inter({ subsets: ['latin'] })



export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <LocalizationProvider dateAdapter={AdapterDayjs}>
      <body>
      
      <Header />
      {children}
      <CssBaseline />
      </body>
      </LocalizationProvider>
    </html>
  )
}
