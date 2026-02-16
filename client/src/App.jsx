export default function App() {
  return (
    // 1. Test Layout & Spacing (flex, min-h, items-center)
    // 2. Test Custom Colors (bg-slate-900)
    <div className="min-h-screen bg-slate-900 flex flex-col items-center justify-center p-4">
      
      <div className="max-w-md w-full bg-slate-800 border border-slate-700 rounded-3xl p-8 shadow-2xl text-center">
        
        {/* 3. Test Typography & Text Gradient (text-5xl, bg-clip-text) */}
        <h1 className="text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-emerald-400 mb-4">
          GoVault
        </h1>

        <p className="text-slate-300 text-lg mb-8">
          If you see a dark background and a gradient title, 
          <span className="text-emerald-400 font-mono"> Tailwind v4 </span> is working perfectly.
        </p>

        {/* 4. Test Interactivity & Hover (hover:scale-105, hover:shadow-blue-500/20) */}
        <button className="w-full py-4 bg-blue-600 text-white font-bold rounded-xl 
                           hover:bg-blue-500 hover:scale-[1.02] active:scale-95 
                           transition-all duration-200 shadow-lg shadow-blue-900/40">
          Verify Connection
        </button>

        {/* 5. Test Grid & Aspect Ratio (grid, aspect-square) */}
        <div className="grid grid-cols-3 gap-2 mt-8">
          {[1, 2, 3].map((i) => (
            <div key={i} className="aspect-square bg-slate-700/50 rounded-lg animate-pulse" />
          ))}
        </div>
      </div>

      <footer className="mt-8 text-slate-500 text-sm">
        Microservices Running on Port 9000
      </footer>
    </div>
  );
}