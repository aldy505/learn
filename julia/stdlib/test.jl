using Test

function calculateNumber(x, y)
  x + y
end

@testset "Random Test" begin
  @testset "Basic maths" begin
    @test calculateNumber(3, 5) == 8
    @test abs(-100) == 100
  end
end