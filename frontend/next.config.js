/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "standalone",
    images: {
        remotePatterns: [
            {
                protocol: "https",
                hostname: "citgen.rcw.io",
                pathname: "**"
            }
        ]
    }
}

module.exports = nextConfig
