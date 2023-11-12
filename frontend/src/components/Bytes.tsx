import { use, useEffect, useMemo, useState } from "react";
import process from "process";
import { resolve } from "path";
import { BytesPrice } from "@/types/Bytes";






interface BytesPriceDisplayProps {
    priceEther: string
    priceUSD: string
}

export default function Bytes({priceEther, priceUSD}: BytesPriceDisplayProps) {
    return (
        <div className="card">
            <div className="card-header">
                 Price
            </div>
            <div className="card-body">
                <h5 className="card-title">${priceUSD} / {priceEther} Ξ</h5>
            </div>
        </div>
    )
} 