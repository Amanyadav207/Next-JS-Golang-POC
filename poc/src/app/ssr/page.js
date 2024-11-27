import React from 'react';
import axios from 'axios';

export default async function SSRPage() {
  try {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL
    console.log("hello", apiUrl);
  
    const res = await fetch(`${apiUrl}/api/ssr-products`, { 
        cache: 'no-store' // Ensures fresh data on each request
      });
      const products = await res.json();

    return (
      <div className="min-h-screen bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-6xl mx-auto">
          <h1 className="text-4xl font-bold text-center text-blue-600 mb-8">
            Server-Side Rendered (SSR) Products
          </h1>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {products.map(product => (
              <div 
                key={product.id} 
                className="bg-white shadow-lg rounded-xl overflow-hidden transform transition-all hover:scale-105 hover:shadow-2xl"
              >
                <div className="relative">
                  <div className="absolute top-4 right-4 z-10">
                    <span className={`px-3 py-1 rounded-full text-xs font-semibold ${
                      product.in_stock 
                        ? 'bg-green-100 text-green-800' 
                        : 'bg-red-100 text-red-800'
                    }`}>
                      {product.in_stock ? 'In Stock' : 'Out of Stock'}
                    </span>
                  </div>
                  <img 
                    src={product.image_url} 
                    alt={product.name} 
                    className="w-full h-48 object-cover"
                  />
                </div>
                <div className="p-6">
                  <div className="flex justify-between items-center mb-3">
                    <h2 className="text-2xl font-semibold text-gray-800">
                      {product.name}
                    </h2>
                    <div className="flex items-center">
                      <svg className="w-5 h-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.286 3.958a1 1 0 00.95.69h4.175c.969 0 1.371 1.24.588 1.81l-3.38 2.458a1 1 0 00-.364 1.118l1.287 3.958c.3.921-.755 1.688-1.54 1.118l-3.38-2.458a1 1 0 00-1.175 0l-3.38 2.458c-.784.57-1.838-.197-1.54-1.118l1.287-3.958a1 1 0 00-.364-1.118L2.49 9.385c-.783-.57-.38-1.81.588-1.81h4.175a1 1 0 00.95-.69l1.286-3.958z" />
                      </svg>
                      <span className="ml-1 text-gray-600">{product.rating}</span>
                    </div>
                  </div>
                  <p className="text-gray-600 mb-4 line-clamp-2">
                    {product.description}
                  </p>
                  <div className="flex justify-between items-center">
                    <span className="text-2xl font-bold text-blue-600">
                      ${product.price.toFixed(2)}
                    </span>
                    <span className="text-sm text-gray-500">
                      {product.brand} | {product.category}
                    </span>
                  </div>
                  <button className="mt-4 w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 transition-colors">
                    Add to Cart
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    );
  } catch (error) {
    console.error('Error fetching SSR products:', error);
    return (
      <div className="min-h-screen bg-red-100 flex items-center justify-center">
        <div className="bg-white p-8 rounded-lg shadow-xl text-center">
          <h2 className="text-3xl text-red-600 mb-4">Error Loading Products</h2>
          <p className="text-gray-700">Unable to fetch product information</p>
        </div>
      </div>
    );
  }
}

// export const dynamic = 'force-dynamic';