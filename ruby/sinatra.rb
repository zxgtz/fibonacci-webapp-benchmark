require 'sinatra'

def fib(n)
	if n == 0
		return 0
	elsif n == 1
		return 1
	else
		return fib(n-1) + fib(n-2)
	end
end

get '/:number' do
	"Ruby + Sinatra<hr> fib(#{params[:number]}): #{fib(params[:number].to_i).to_s}"
end
