import { use, useEffect, useMemo, useState } from "react";
import process from "process";
import { resolve } from "path";
import { BytesPrice } from "@/types/Bytes";



export default function Bytes({priceEther, priceUSD}: BytesPrice) {
    return (
        <div>
            <span>Bytes Price: {priceEther} Ξ / ${priceUSD}</span>
        </div>
    )
} 