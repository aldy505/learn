using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using Microsoft.EntityFrameworkCore;

namespace dadjokes {

  public class Joke : BaseModel {
    [Required, StringLength(250)]
    public string JokeText { get; set; }
  }

  public class BaseModel {
    [Key, DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public long ID { get; set; }
  }
  
  public class AppContext : DbContext {
    public DbSet<Joke> Joke { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder options) {
      options.UseSqlite("Data Source=DadJokes.db");
    }
  }
}