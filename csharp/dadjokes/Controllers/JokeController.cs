using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace dadjokes.Controllers
{
    [ApiController]
    [Route("joke")]
    public class JokeController : ControllerBase
    {
    private const string V = "{id:int}";
    private const string V1 = "application/json";
    private readonly AppContext _context;
        public JokeController() {
            _context = new AppContext();
        }

        [HttpGet]
        public dynamic GetAll() {
            var jokes = _context.Joke.ToList();
            return jokes;
        }

        [HttpGet("{id:long}")]
        public async Task<dynamic> GetByID([FromRoute] long id) {
            var joke = await _context.Joke.FindAsync(id);
            return joke;
        }

        [HttpPut]
        [Consumes(V1)]
        public async Task<dynamic> Add([FromBody] Joke joke) {
            _context.Joke.Add(joke);
            await _context.SaveChangesAsync();

            return joke;
        }

        [HttpPatch("{id:long}")]
        [Consumes(V1)]
        public async Task<dynamic> Edit([FromRoute] long id, [FromBody] Joke joke) {
            // Asynchronously finds an entity with the given primary key values. 
            // If an entity with the given primary key values exists in the context, 
            // then it is returned immediately without making a request to the store. 
            // Otherwise, a request is made to the store for an entity with the given 
            // primary key values and this entity, if found, is attached to the context 
            // and returned. If no entity is found in the context or the store, then null is returned.
            var jokeItem = await _context.Joke.FindAsync(id);
            if (jokeItem == null) {
                return NotFound();
            }

            _context.Entry(jokeItem).CurrentValues.SetValues(new Dictionary<string, string>() {
                { "JokeText", joke.JokeText }
            });
            await _context.SaveChangesAsync();

            return new Dictionary<string, dynamic>() {
                { "id", id },
                { "jokeText", joke.JokeText }
            };
        }

        [HttpDelete("{id:long}")]
        public async Task<dynamic> Delete([FromRoute] long id) {
            var jokeItem = await _context.Joke.FindAsync(id);
            if (jokeItem == null) {
                return NotFound();
            }

            _context.Joke.Remove(jokeItem);
            await _context.SaveChangesAsync();

            return jokeItem;
        }
    }
}
