export const category = [
    {
        icon: "/img/category_icon/mac_shop.svg",
        name: "Mac",
        slug: "/mac/"
    },
    {
        icon: "/img/category_icon/shop_iphone.svg",
        name: "iPhone",
        slug: "/iphone/",
    },
    {
        icon: "/img/category_icon/shop_ipad.svg",
        name: "iPad",
        slug: "/ipad/"
    },
    {
        icon: "/img/category_icon/airpods.svg",
        name: "AirPods",
        slug: "/airpods/"
    },
    {
        icon: "/img/category_icon/watch_nav_shop_watch.svg",
        name: "Watch",
        slug: "/watch/"
    },
    {
        icon: "/img/category_icon/apple_tv.svg",
        name: "TV & Home",
        slug: "/tv-and-home/"
    },
    {
        icon: "/img/category_icon/home_accessories_light.svg",
        name: "Accessories",
        slug: "/accessories/",
        _id: "111",
        subCat: [
            {
                icon: "/img/category_icon/mac_shop.svg",
                name: "Mac",
                slug: "/mac/"
            },
            {
                icon: "/img/category_icon/shop_iphone.svg",
                name: "iPhone",
                slug: "/iphone/",
            },
            {
                icon: "/img/category_icon/shop_ipad.svg",
                name: "iPad",
                slug: "/ipad/"
            },
            {
                icon: "/img/category_icon/airpods.svg",
                name: "AirPods",
                slug: "/airpods/"
            },
            {
                icon: "/img/category_icon/watch_nav_shop_watch.svg",
                name: "Watch",
                slug: "/watch/"
            },
            {
                icon: "/img/category_icon/apple_tv.svg",
                name: "TV & Home",
                slug: "/tv-and-home/"
            },
        ]
    },
    {
        icon: "/img/category_icon/support.svg",
        name: "iSupport",
        slug: "/admin/support/"
    },
    {
        icon: "/img/category_icon/admin.svg",
        name: "Admin",
        slug: "/admin/dashboard/"
    },
]

export const products = [
    {
        "brand": "Apple",
        "category": "Mac",
        "slug": "mac",
        "name": "Apple MacBook Pro",
        "model": "MLVF3H",
        "price": {
            "currencyIso": "INR",
            "discount": 19910,
            "mrp": 239900,
            "sellingPrice": 219990,
        },
        "images": "https://media.croma.com/image/upload/v1647427408/Croma%20Assets/Computers%20Peripherals/Laptop/Images/245235_hcqtej.png",
        "color": ['#2e3642', '#f0e5d2', '#7d7e80', '#e3e4e6']
    },
    {
        "brand": "Apple",
        "category": "iPhone",
        "slug": "iphone/iphone-13-pro/MLVF3H",
        "name": "Apple iPhone 13 Pro",
        "model": "MLVF3H",
        "price": {
            "currencyIso": "INR",
            "discount": 5910,
            "mrp": 129900,
            "sellingPrice": 123990,
        },
        "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
        "color": ['#566856', '#f1f3ed', '#f9e8d2', '#605f5d', '#b2c9dd']
    },
    {
        "brand": "Apple",
        "category": "iPad",
        "slug": "ipad",
        "name": "Apple iPad Pro",
        "model": "MHNG3HN/A",
        "price": {
            "currencyIso": "INR",
            "discount": 19910,
            "mrp": 239900,
            "sellingPrice": 219990,
        },
        "images": "https://media.croma.com/image/upload/v1619272147/Croma%20Assets/Computers%20Peripherals/Tablets%20and%20iPads/Images/234301_9_vxz8o4.png",
        "color": ['#eaebed', '#7f8084']
    },
    {
        "brand": "Apple",
        "category": "Watch",
        "slug": "watch",
        "name": "Apple Watch Series 7",
        "model": "MKN53HN/A",
        "price": {
            "currencyIso": "INR",
            "discount": 3000,
            "mrp": 44900,
            "sellingPrice": 41900,
        },
        "images": "https://media.croma.com/image/upload/v1633696522/Croma%20Assets/Wearable/Wearable%20Devices/Images/244886_ipofha.png",
        "color": ['#1a2531', '#ded6d1', '#35382c', '#6d8dba', '#aa393b', '#3e3a35', '#e6e6e6', '#e5caad', '#4c4845', '#d1cdca']
    },
    {
        "brand": "Apple",
        "category": "AirPods",
        "slug": "airpods",
        "name": "Apple AirPods Pro ",
        "model": "MLWK3HN/A",
        "price": {
            "currencyIso": "INR",
            "discount": 4801,
            "mrp": 26300,
            "sellingPrice": 21499,
        },
        "images": "https://media.croma.com/image/upload/v1651838583/Croma%20Assets/Entertainment/Wireless%20Earbuds/Images/251322_qd8g3s.png",
    }
]

export const product = {
    "category": "iPhone",
    "name": "Apple iPhone 14",
    "base_price" : 79900,
    "specifications": [
        {"name": "Display",  "description" : "11 Inches 120 Hz Liquid Retina Display"},
        {"name": "Capacity", "description" :  "128GB, 256GB, 512GB, 1TB, 2TB"},
        {"name": "Chip", "description" :  "Apple M1 Chip"},
        {"name": "Front Camera", "description" :  "12 MP : TrueDepth Camera, ƒ/2.4, 122 Degree Field of View, Animoji and Memoji, Lens Correction"},
        {"name": "Rear Camera", "description" :  "12 MP : Wide Camera, ƒ/1.8, Five‑Element Lens, Autofocus with Focus Pixels | 10 MP : Ultra-Wide Camera, ƒ/2.4, 125 Degree Field of View, Five‑Element Lens, Lens Correction | 12MP Telephoto Camera: ƒ/2.8 Aperture"},
        {"name": "Sound",  "description" : "Quad Speakers | AAC‑LC, HE‑AAC, HE‑AAC v2, Protected AAC, MP3, Linear PCM, Apple Lossless, FLAC, Dolby Digital (AC‑3), Dolby Digital Plus (E‑AC‑3), Dolby Atmos and Audible (formats 2, 3, 4, Audible Enhanced Audio, AAX and AAX+)"},
        {"name": "WiFi", "description" :  "IEEE 802.11a/b/g/n/ac/ax | Dual Band (2.4 GHz and 5 GHz), HT80 with MIMO"},
        {"name": "Bluetooth", "description" :  "Bluetooth 5.0"},
        {"name": "Battery Life", "description" :  "Up to 10 Hours of Surfing the Web on Wi-Fi or Watching Video, Built-in 28.65 WHr Rechargeable Battery"},
        {"name": "Connector", "description" :  "USB Type-C | 20 Watts Fast Charging, Charging Via Power Adapter or USB-C to Computer System"},
        {"name": "Dimensions", "description" :  "7.02 x 0.23 x 9.74 Inches"},
        {"name": "Weight", "description" :  "466 g"},
        {"name": "In the Box", "description" :  "1 U Tablet | Power Adapter | USB Cable | Quick Start Guide | Warranty Card"},
        {"name": "Warranty", "description" :  "12 Months | International Travellers | Carry-In"}
    ],
}

export const productImages = [
    "https://media.croma.com/image/upload/v1631776566/Croma%20Assets/Communication/Mobiles/Images/243463_wwwi8s.png",
    "https://media.croma.com/image/upload/v1631776563/Croma%20Assets/Communication/Mobiles/Images/243463_3_pau5he.png",
    "https://media.croma.com/image/upload/v1631776562/Croma%20Assets/Communication/Mobiles/Images/243463_6_qepfxn.png",
    "https://media.croma.com/image/upload/v1631776565/Croma%20Assets/Communication/Mobiles/Images/243463_7_zzy5gw.png",
]

export const productAccessories = [
    {
        "model": "MLVF3A",
        "name": "Apple MagSafe 15 Watts Wireless Charger",
        "sellingPrice": 4500,
        "slug": "/accessories/apple-magsafe-15-watts-wireless-charger/MHXH3ZMA"
    },
    {
        "model": "MLVF3B",
        "name": "Apple 20 Watt USB-C Power Adapter",
        "sellingPrice": 1900,
        "slug": "/accessories/apple-20-watt-usb-c-power-adapter/MHXH3ZMA"
    },
    {
        "model": "MLVF3C",
        "name": "Apple PVC 1 Meter USB 3.0 (Type-C) to Lightning Charging Cable",
        "sellingPrice": 1900,
        "slug": "/accessories/apple-pvc-1-meter-usb-3-0-type-c-to-lightning-charging-cable/MHXH3ZMA"
    },
]

export const relatedProducts = [
    {
        "brand": "Apple",
        "category": "iPhone",
        "slug": "iphone/iphone-13-pro/MLVF3H",
        "name": "Apple iPhone 13 Pro",
        "model": "MLVF3H",
        "price": {
            "currencyIso": "INR",
            "discount": 5910,
            "mrp": 129900,
            "sellingPrice": 123990,
        },
        "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
        "color": ['#566856', '#f1f3ed', '#f9e8d2', '#605f5d', '#b2c9dd']
    }, {
        "brand": "Apple",
        "category": "iPhone",
        "slug": "iphone/iphone-13-pro/MLVF3H",
        "name": "Apple iPhone 13 Pro",
        "model": "MLVF3H",
        "price": {
            "currencyIso": "INR",
            "discount": 5910,
            "mrp": 129900,
            "sellingPrice": 123990,
        },
        "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
        "color": ['#566856', '#f1f3ed', '#f9e8d2', '#605f5d', '#b2c9dd']
    }
]

export const orders = [
    {
        "orderNumber" : "403-4481972-2195557",

        "category": "iPhone",
        "name": "Apple iPhone 13 Pro",
        "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
        "shipping" : 0,
        "tax" : 999,
        "total" : 99999,

        "shippingAddress" : {
            "firstname" : "Abhinay Narayan",
            "lastname" : "Singh",
            "address" : "Aashirvad Niketan, Seer Goverdhanpur, Near Chaura Mata Mandir, Banaras Hindu University",
            "city" : "Varanasi",
            "state" : "Uttar Pradesh",
            "country" : "India",
            "postalCode" : "221005",
            "phoneNumber": "8368380854"
        },
        "billingAddress" : {
            "firstname" : "Abhinay Narayan",
            "lastname" : "Singh",
            "address" : "Aashirvad Niketan, Seer Goverdhanpur, Near Chaura Mata Mandir, Banaras Hindu University",
            "city" : "Varanasi",
            "state" : "Uttar Pradesh",
            "country" : "India",
            "postalCode" : "221005",
            "phoneNumber": "8368380854"
        },
        
        "paymentMethod" : "BHIM UPI",
        "rewardPoint" : 100,

        "carrier" : {
            "expectedDeliveryDate" : "Wed 17 Aug"
        }
    }
]

export const order = {
        "orderNumber" : "403-4481972-2195557",

        "category": "iPhone",
        "name": "Apple iPhone 13 Pro",
        "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
        "shipping" : 0,
        "tax" : 999,
        "total" : 144998,
        "isBillingAddressIsSame" : false,
        "reward" : 100,
        "isReturnWindowClosed" : false,
        "item" : [
            {
                "name": "Apple iPhone 13 Pro",
                "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
                "price" : 124999,
                "varient" : {
                    "model" : "iPhone 13 Pro",
                    "finish" : "Green",
                    "storage" : "256 GB"
                }
            },
            {
                "name": "Apple AirPods Pro ",
                "images": "https://media.croma.com/image/upload/v1651838583/Croma%20Assets/Entertainment/Wireless%20Earbuds/Images/251322_qd8g3s.png",
                "price" : 19999
            },
        ],
        
        "shippingAddress" : {
            "firstname" : "Abhinay Narayan",
            "lastname" : "Singh",
            "address" : "Aashirvad Niketan, Seer Goverdhanpur, Near Chaura Mata Mandir, Banaras Hindu University",
            "city" : "Varanasi",
            "state" : "Uttar Pradesh",
            "country" : "India",
            "postalCode" : "221005",
            "phoneNumber": "8368380854"
        },
        "billingAddress" : {
            "firstname" : "Abhinay Narayan",
            "lastname" : "Singh",
            "address" : "Aashirvad Niketan, Seer Goverdhanpur, Near Chaura Mata Mandir, Banaras Hindu University",
            "city" : "Varanasi",
            "state" : "Uttar Pradesh",
            "country" : "India",
            "postalCode" : "221005",
            "phoneNumber": "8368380854"
        },        
        "paymentMethod" : "BHIM UPI",
        "rewardPoint" : 100,
        
        "carrier" : {
            expectedDeliveryDate: 1668849812000,
            carrierURL: "http://www.xpressbees.com/",
            carrierLogo: "https://xinglian-dev-1254213275.cos.ap-guangzhou.myqcloud.com/provider/xpressbees.gif?sign=q-sign-algorithm%3Dsha1%26q-ak%3DAKIDObYFfogNnwoOimZ5KIzOYREk01lDOJyN%26q-sign-time%3D1669284364%3B1700820364%26q-key-time%3D1669284364%3B1700820364%26q-header-list%3Dhost%26q-url-param-list%3D%26q-signature%3Dd8385b480120b68ab095f3b4a9a984141a646057",
        },
        trackingDetails: [
            {
              date: 1668581626000,
              city: "AGRA",
              statusDetails: "Shipment Created"
            },
            {
              date: 1668599670000,
              city: "AGRA",
              statusDetails: "Pickup From Seller"
            },
            {
              date: 1668657805000,
              city: "AGRA",
              statusDetails: "Received at FKL_farukngr"
            },
            {
              date: 1668581647000,
              city: "Gurgaon",
              statusDetails: "Dispatched to Motherhub_JKS"
            },
            {
              date: 1668658462000,
              city: "Farukhnagar",
              statusDetails: "Received at MotherHub_FRK"
            },
            {
              date: 1668697466000,
              city: "Varanasi",
              statusDetails: "Dispatched to Motherhub_VNI"
            },
            {
              date: 1668773966000,
              city: "Varanasi",
              statusDetails: "Received at Motherhub_VNI"
            },
            {
              date: 1668773967000,
              city: "Varanasi",
              statusDetails: "Received at Motherhub_VNI"
            },
            {
              date: 1668801026000,
              city: "Varanasi",
              statusDetails: "Dispatched to BHUHub_VNS"
            },
            {
              date: 1668822588000,
              city: "Varanasi",
              statusDetails: "Received at BHUHub_VNS"
            },
            {
              date: 1668827546000,
              city: "VARANASI",
              statusDetails: "Dispatched to Ajit Genaral store"
            },
            {
              date: 1668832597000,
              city: "VARANASI",
              statusDetails: "Received at Ajit Genaral store"
            },
            {
              date: 1668832741000,
              city: "VARANASI",
              statusDetails: "Out For Delivery"
            }
          ],
        "trackingStatus" : [
            {
              "name" : "Order Placed",
              "timeStamp" : 1668581626000,
            },
            {
              "name" : "Order Confirmed",
              "timeStamp" : 1668581626000,
            },
            {
              "name" : "In Transit",
              "timeStamp" : 1668581626000,
            },
            {
              "name" : "Out For Delivery",
              "timeStamp" : 1668827546000,
            },
            {
              "name" : "Delivered",
            },
          ]
    }


export const bag = {
    bagTotal: 145489,
    bagSubTotal: 123295,
    bagSubTotalTax: 22194,
    shipping: "FREE",
    items: [

        {
            "brand": "Apple",
            "category": "iPhone",
            "slug": "iphone/iphone-13-pro/MLVF3H",
            "name": "Apple iPhone 13 Pro 512GB Product RED",
            "model": "MLVF3H",
            "price": {
                "currencyIso": "INR",
                "discount": 5910,
                "mrp": 129900,
                "sellingPrice": 123990,
            },
            "images": "https://media.croma.com/image/upload/v1631855852/Croma%20Assets/Communication/Mobiles/Images/243512_7_ulnxdo.png",
        },
        {
            "brand": "Apple",
            "category": "AirPods",
            "slug": "airpods",
            "name": "Apple AirPods Pro",
            "model": "MLWK3HN/A",
            "price": {
                "currencyIso": "INR",
                "discount": 4801,
                "mrp": 26300,
                "sellingPrice": 21499,
            },
            "images": "https://media.croma.com/image/upload/v1651838583/Croma%20Assets/Entertainment/Wireless%20Earbuds/Images/251322_qd8g3s.png",
        },
        {
            "brand": "Apple",
            "category": "Mac",
            "slug": "mac",
            "name": "Apple MacBook Pro",
            "model": "MLVF3H",
            "price": {
                "currencyIso": "INR",
                "discount": 19910,
                "mrp": 239900,
                "sellingPrice": 219990,
            },
            "images": "https://media.croma.com/image/upload/v1647427408/Croma%20Assets/Computers%20Peripherals/Laptop/Images/245235_hcqtej.png",
            "color": ['#2e3642', '#f0e5d2', '#7d7e80', '#e3e4e6']
        },
    ]
}

export const user = {
    ID: "6330182d7ec86b5b0da8722c",
    FirstName: "Abhinay Narayan",
    LastName: "Singh",
    Email: "abhinay@admin.com",
    Phone: "8795675599",
    IsActive: true,
    IsAdmin: true,
    Created_at: "2022-09-25T08:58:22Z",
    Updated_at: "2022-09-25T08:58:22Z",
    Last_login: "2022-09-28T13:04:08Z",
    Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjYzMzAxODJkN2VjODZiNWIwZGE4NzIyYyIsIkZpcnN0TmFtZSI6IkFiaGluYXkiLCJMYXN0TmFtZSI6IlNpbmdoIiwiRW1haWwiOiJhYmhpbmF5QGFkbWluLmNvbSIsIlBob25lIjoiODc5NTY3NTU5OSIsIklzQWRtaW4iOnRydWUsIklzQWN0aXZlIjp0cnVlLCJleHAiOjE2NjQ1MzM5NjN9.uwXnV1AiMTbXovECEdemBeoKXRL77_p5kp4MR0fqN8g"
  }

export const varient = {
    "storage": [
            {  
                "name" : "128GB",
                "description" : "",
                "addonPrice" : 0
            },
            {  
                "name" : "256GB",
                "description" : "",
                "addonPrice" : 10000
            },
            {  
                "name" : "512GB",
                "description" : "",
                "addonPrice" : 20000
            },
        ],
    "finish": [
        {
            "addonPrice" : 0,
            "name": "Blue",
            "additionalDetails": "#286788",
        },
        {
            "addonPrice" : 0,
            "name": "Pink",
            "additionalDetails": "#f8ddd7",
        },
        {
            "addonPrice" : 10000,
            "name": "Green",
            "additionalDetails": "#3a4c37",
        },
        {
            "addonPrice" : 0,
            "name": "Midnight",
            "additionalDetails": "#232a31",
        },
        {
            "addonPrice" : 10000,
            "name": "Starlight",
            "additionalDetails": "#faf6f2",
        },
        {
            "addonPrice" : 0,
            "name": "Product RED",
            "additionalDetails": "#c23e16",
        },
    ],
    "model": [
        {"name": "iPhone 14", "description" : "15.4 cm (6.1-inch) display", "addonPrice" : 79900},
        {"name": "iPhone 14 Plus", "description" : "13.7 cm (5.4-inch) display", "addonPrice" : 89900}, 
    ],
    "memory": [
        {
            "name": "4GB",
            "addonPrice": 0
        },
        {
            "name": "6GB",
            "addonPrice": 10000
        }
    ],
    "processor" : [
        {
            "name" : "A15 Bionic",
            "addonPrice" : 0
        },
        {
            "name" : "A16 Bionic",
            "addonPrice" : 10000
        }
    ],
    "connectivity" : [
        {
            "name" : "5G",
            "addonPrice" : 0
        },
        {
            "name" : "6G",
            "addonPrice" : 10000
        }
    ],
}

export const hero = []