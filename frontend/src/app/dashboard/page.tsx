"use client";
import { dashboard } from '@/models/user';
import { useEffect, useState } from 'react';

export default function Dashboard() {
    const [userData, setUserData] = useState<dashboard | null>(null);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchData = async () => {
            try {
                const token = localStorage.getItem('authToken');
                const userId = localStorage.getItem('userId');

                if (!token || !userId) {
                    throw new Error('User not authenticated');
                }
                const response = await fetch('http://localhost:3001/api/get', {
                    method: 'get',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': "Bearer " + token,
                        'Origin': 'http://localhost:3000',
                    },
                });
               
                if (!response.ok) {
                    const errorText = await response.text();
                    console.error('Raw server error:', errorText);
                    throw new Error(`Failed to fetch user data: ${errorText}`);
                }

                const data = await response.json();
                console.log('Response:', data);
                setUserData(data);
            } catch (err: any) {
                console.error('Error fetching user data:', err);
                setError(err.message);
            }
        };

        fetchData();
    }, []); // Run only once when the component mounts

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
            <div className="max-w-md w-full space-y-8 p-10 bg-white rounded-lg shadow-md">
                <div>
                    <h1 className="text-center text-3xl font-extrabold text-gray-900">
                        Welcome User: {userData ? userData["username"] : 'User'}
                    </h1>
                </div>
                {error && (
                    <div className="text-red-500 text-center">
                        {error}
                    </div>
                )}
            </div>
        </div>
    );
}