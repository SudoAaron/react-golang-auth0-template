import React from "react";
import loading from '../../assets/loading.svg'
import Image from "next/image";

export default function LoadingSpinner() {
    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <div className="spinner">
        <div className="w-16 h-16 rounded-full">
            <Image priority src={loading} alt="Loading" className="w-16 h-16 mx-auto my-auto z-50" />
        </div>
        </div>
        </main>
    )
}